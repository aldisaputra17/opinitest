# Go Jwt Implementation & Middleware Gin-Gonic

##Overview
*CRUD User, 
Postingan
*Authenication Login, 
Register, 
Middleware JWT

## Endpoint
*GroupAuth("/api/auth")
  authRoutes.POST("/login", authController.Login)
  authRoutes.POST("/register", authController.Register)
  
*GroupUser("/api/user")
  userRoutes.GET("/profile", userController.Profile)
  userRoutes.PUT("/profile", userController.Update)
	
*GroupPostingan("/api/v1")
  postinganRoutes.GET("/posts", postinganController.GetAll)
  postinganRoutes.GET("/post/:id", postinganController.FindByID)
  postinganRoutes.POST("/post", postinganController.Create)
  postinganRoutes.PUT("/post/:id", postinganController.Update)
  postinganRoutes.DELETE("/post/:id", postinganController.Delete)

##Tools
*Mysql
*GORM
*GIN-GONIC
*JWT token

