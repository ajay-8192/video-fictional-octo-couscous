package middlewares

import "github.com/gin-gonic/gin"

func Authentication(ctx *gin.Context) {
	ctx.Next();
}
