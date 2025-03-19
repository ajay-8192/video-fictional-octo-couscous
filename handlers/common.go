package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
