package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/aldisaputra17/post_opinia/dataobject"
	"github.com/aldisaputra17/post_opinia/helpers"
	"github.com/aldisaputra17/post_opinia/models"
	"github.com/aldisaputra17/post_opinia/service"
	"github.com/gin-gonic/gin"
)

type PostinganController interface {
	GetAll(context *gin.Context)
	FindByID(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type postController struct {
	postService service.PostinganService
	jwtService  service.JWTService
}

func NewPostController(postServ service.PostinganService, jwtServ service.JWTService) PostinganController {
	return &postController{
		postService: postServ,
		jwtService:  jwtServ,
	}
}

func (c *postController) GetAll(context *gin.Context) {
	var posts []models.Postingan = c.postService.GetAll()
	res := helpers.BuildResponse(true, "OK", posts)
	context.JSON(http.StatusOK, res)
}

func (c *postController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})

	var post models.Postingan = c.postService.FindById(id)
	if (post == models.Postingan{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", post)
		context.JSON(http.StatusOK, res)
	}
}

func (c *postController) Create(context *gin.Context) {
	var postCreateObj dataobject.PostCreateObj
	errObj := context.ShouldBind(&postCreateObj)
	if errObj != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			postCreateObj.UserID = convertedUserID
		}
		result := c.postService.Create(postCreateObj)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)

	}
}

func (c *postController) Update(context *gin.Context) {
	var postUpdatedObj dataobject.PostUpdatedObj
	errObj := context.ShouldBind(&postUpdatedObj)
	if errObj != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.postService.IsAllowedToEdit(userID, postUpdatedObj.UserID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			postUpdatedObj.UserID = id
		}
		result := c.postService.Update(postUpdatedObj)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)

	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *postController) Delete(context *gin.Context) {
	var post models.Postingan
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	post.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.postService.IsAllowedToEdit(userID, post.ID) {
		c.postService.Delete(post)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *postController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
