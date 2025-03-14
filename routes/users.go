package routes

import (
	"chaterminal/handlers"
	"chaterminal/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine) {
    user := router.Group("/user")
    {
        user.POST("/create", handlers.RegisterUser)
        user.POST("/send-otp", handlers.SendOTPForEmailRegister)
        user.POST("/verify-otp", handlers.VerifyEmailOtpForRegister)
        user.GET("/details", middlewares.Authentication, handlers.GetUserDetails)
        user.PUT("/update", middlewares.Authentication, handlers.UpdateUserDetails)
        user.POST("/deactivate", middlewares.Authentication, handlers.DeactivateUser)
    }
}
