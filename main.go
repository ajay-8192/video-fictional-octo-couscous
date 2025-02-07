package main

import (
	"chaterminal/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	
	fmt.Println("Hello World!!")
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to chatter box",
		})
	})

	routes.Routes(router)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success",
		})
	})

	router.Run(":3000")
}
