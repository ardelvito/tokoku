package routes

import (
    "github.com/gin-gonic/gin"
    "backendmono/controllers"
    "backendmono/middlewares"
)

func SetupRouter(userController *controllers.UserController, jwtSecret string) *gin.Engine {
    router := gin.Default()

    router.POST("/login", userController.Login)
    router.POST("/register", userController.Register)

    authorized := router.Group("/")
    authorized.Use(middlewares.JWTAuthMiddleware(jwtSecret))
    {
        authorized.POST("/changepassword", userController.ForgotPassword)
        authorized.POST("/editprofile", userController.EditProfile)
        authorized.GET("/profile/:id", userController.GetProfile)
    }

    return router
}
