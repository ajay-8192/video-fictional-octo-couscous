package routes

import (
	"chaterminal/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/status", handlers.SuccessMessage)

	api := router.Group("/api")
	{
		api.GET("/status", handlers.SuccessMessage)
	}
	
	frontendRoutes(router)
	userRoutes(router)
}
