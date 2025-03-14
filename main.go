package main

import (
	"chaterminal/routes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))

	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to chatter box",
		})
	})

	routes.Routes(router)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "PONG",
		})
	})

	router.Run(":3000")
}
