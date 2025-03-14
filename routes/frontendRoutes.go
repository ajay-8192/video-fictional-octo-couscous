package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func frontendRoutes(router *gin.Engine) {
	uiroutes := router.Group("/")
	{
		uiroutes.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Chaterminal",
				"message": "Welcome to Chaterminal",
			})
		})
	}
}
