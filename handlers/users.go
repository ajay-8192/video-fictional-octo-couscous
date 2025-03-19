package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {

	var requestBody struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Number    string `json:"number"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User registered",
	})
}

func SendOTPForEmailRegister(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to email",
	})
}

func VerifyEmailOtpForRegister(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP verified successfully",
	})
}

func GoogleSSOLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Google SSO login successful",
	})
}

func GetUserDetails(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User details fetched Successfully",
	})
}

func UpdateUserDetails(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User details updated successfully",
	})
}

func DeactivateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deactivated successfully",
	})
}
