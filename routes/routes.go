package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {



	api := router.Group("/api")
	{
		api.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	frontendRoutes(router)
	userRoutes(router)
}
