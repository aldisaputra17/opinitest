package main

import (
	"github.com/aldisaputra17/post_opinia/controllers"
	"github.com/aldisaputra17/post_opinia/database"
	"github.com/aldisaputra17/post_opinia/repository"
	"github.com/aldisaputra17/post_opinia/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                        = database.DBConnection()
	userRepository      repository.UserRepository       = repository.NewUserRepository(db)
	postinganRepository repository.PostinganRepository  = repository.NewPostinganRepository(db)
	jwtService          service.JWTService              = service.NewJWTService()
	userService         service.UserService             = service.NewUserService(userRepository)
	postinganService    service.PostinganService        = service.NewPostService(postinganRepository)
	authService         service.AuthService             = service.NewAuthService(userRepository)
	authController      controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	userController      controllers.UserController      = controllers.NewUserController(userService, jwtService)
	postinganController controllers.PostinganController = controllers.NewPostController(postinganService, jwtService)
)

func main() {
	defer database.CloseDatabaseConnection(db)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("/api/user")

	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	postinganRoutes := r.Group("/api/v1")

	{
		postinganRoutes.GET("/posts", postinganController.GetAll)
		postinganRoutes.GET("/post/:id", postinganController.FindByID)
		postinganRoutes.POST("/post", postinganController.Create)
		postinganRoutes.PUT("/post/:id", postinganController.Update)
		postinganRoutes.DELETE("/post/:id", postinganController.Delete)
	}
	r.Run()
}
