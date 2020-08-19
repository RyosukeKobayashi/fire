package router

import (
	"api/controller"
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	router.POST("/createReset", controller.InitiatePasswordReset)
	router.POST("/resetPassword", controller.ResetPassword)
	return router
}
