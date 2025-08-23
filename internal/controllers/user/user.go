package controller

import (
	"net/http"
	"workout_tracker/internal/config"
	model "workout_tracker/internal/model/user"
	"workout_tracker/pkg/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ChangePassword struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// @Tags User
// @Summary Get user profile
// @Description Get the profile information of the authenticated user
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetMyProfile(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user model.User
	if err := config.GetDB().Where("ID = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User profile retrieved successfully", "data": gin.H{
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"email":       user.Email,
		"is_verified": user.IsVerified,
	}})
}

// @Tags User
// @Summary Change user password
// @Description Change the password of the authenticated user
// @Param request body ChangePassword true "Change Password Request"
// @Accept json
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/change-password [patch]
func UpdatePassword(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var reqBody ChangePassword
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if reqBody.NewPassword != reqBody.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password and confirm password do not match"})
		return
	}

	var user model.User
	if err := config.GetDB().Where("ID = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	verifyPassword, err := argon2id.ComparePasswordAndHash(reqBody.OldPassword, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error verifying password"})
		return
	}
	if !verifyPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid old password"})
		return
	}

	hash, err := argon2id.CreateHash(reqBody.NewPassword, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing new password"})
		return
	}

	user.Password = hash
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Password updated successfully"})
}
