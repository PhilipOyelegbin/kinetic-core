package routes

import (
	auth "workout_tracker/internal/controllers/auth"
	user "workout_tracker/internal/controllers/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	// auth routes
	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)
	api.POST("/send", auth.SendVerificationEmail)
	api.GET("/verify-email", auth.VerifyEmail)
	api.POST("/forgot-password", auth.SendForgotPasswordEmail)
	api.POST("/reset-password", auth.ResetPassword)

	// user routes
	api.GET("/users/me", user.GetMyProfile)
	api.PATCH("/users/change-password", user.UpdatePassword)
}
