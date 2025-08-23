package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"workout_tracker/internal/config"
	model "workout_tracker/internal/model/user"
	"workout_tracker/pkg/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type RegisterUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PasswordReset struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// @Tags Auth
// @Summary Register a new user
// @Description Register a new user
// @Accept json
// @Produce json
// @Param user body RegisterUser true "User"
// @Success 201 {object} RegisterUser
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	var reqBody model.User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reqBody.FirstName == "" || reqBody.LastName == "" || reqBody.Email == "" || reqBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	if !config.GetDB().Where("email = ?", reqBody.Email).First(&model.User{}).RecordNotFound() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hash, err := argon2id.CreateHash(reqBody.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	reqBody.Password = hash
	token, exp := utils.GenerateCodeAndTime()
	reqBody.VerifyToken = token
	reqBody.VerifyExpTime = exp
	res := config.GetDB().Create(&reqBody)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	subject := "Please verify your email"
	message := fmt.Sprintf("Subject: %s\n\nHi %s,\n\nPlease verify your email by clicking on the following link:\n\n- %s/verify-email?token=%s\n\nThank you!\n\nWarm regards,\n\nKinetic Core Team", subject, reqBody.FirstName, os.Getenv("APP_URL"), reqBody.VerifyToken)
	err = utils.SendEmail(reqBody.Email, reqBody.FirstName, subject, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Verification mail sent, check your junk or promotion folder!"})
}

// @Tags Auth
// @Summary Send verification mail
// @Description Send a verification email to the user
// @Accept json
// @Produce json
// @Param email query string true "User Email"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /send [post]
func SendVerificationEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	var user model.User
	getUser := config.GetDB().Where("email = ?", email).First(&user)
	if getUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if user.IsVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already verified"})
		return
	}

	token, exp := utils.GenerateCodeAndTime()
	user.VerifyToken = token
	user.VerifyExpTime = exp
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save verification details to database"})
		return
	}

	subject := "Please verify your email"
	message := fmt.Sprintf("Subject: %s\n\nHi %s,\n\nPlease verify your email by clicking on the following link:\n\n- %s/verify-email?token=%s\n\nThank you!\n\nWarm regards,\n\nKinetic Core Team", subject, user.FirstName, os.Getenv("APP_URL"), user.VerifyToken)
	err := utils.SendEmail(user.Email, user.FirstName, subject, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mail sent, check your junk or promotion folder!"})
}

// @Tags Auth
// @Summary Verify user email
// @Description Verify a user's email address
// @Param token query string true "Verification Token"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /verify-email [get]
func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
	}

	var user model.User
	if err := config.GetDB().Where("verify_token = ?", token).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	convertedExpTime := time.Unix(user.VerifyExpTime, 0)
	if convertedExpTime.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return
	}

	user.IsVerified = true
	user.VerifyToken = ""
	user.VerifyExpTime = 0
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

// @Tags Auth
// @Summary Login as a user
// @Description Login as a user
// @Accept json
// @Produce json
// @Param user body LoginUser true "Auth"
// @Success 200 {object} LoginUser
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var reqBody model.User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reqBody.Email == "" || reqBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	var user model.User
	verifyEmail := config.GetDB().Where("email = ?", reqBody.Email).First(&user)
	if verifyEmail.Error != nil {
		if verifyEmail.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		log.Printf("Database error during authentication: %v", verifyEmail.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(reqBody.Password, user.Password)
	if err != nil {
		log.Printf("Error during password comparison: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if !user.IsVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not verified"})
		return
	}

	token, err := utils.SignJWTToken(int64(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

// @Tags Auth
// @Summary Send a forgot password mail
// @Description Send a forgot password email to the user
// @Accept json
// @Produce json
// @Param email query string true "User Email"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /forgot-password [post]
func SendForgotPasswordEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	var user model.User
	getUser := config.GetDB().Where("email = ?", email).First(&user)
	if getUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if !user.IsVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not verified"})
		return
	}

	token, exp := utils.GenerateCodeAndTime()
	user.ResetToken = token
	user.ResetExpTime = exp
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reset details to database."})
		return
	}

	subject := "Reset your password"
	message := fmt.Sprintf("Subject: %s\n\nHi %s,\n\nPlease reset your password by clicking on the following link:\n\n- %s/reset-password?token=%s\n\nThank you!\n\nWarm regards,\n\nKinetic Core Team", subject, user.FirstName, os.Getenv("APP_URL"), user.ResetToken)
	err := utils.SendEmail(user.Email, user.FirstName, subject, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send forgot password email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mail sent, check your junk or promotion folder!"})
}

// @Tags Auth
// @Summary Reset user password
// @Description Reset a user's password
// @Param token query string true "Reset Token"
// @Param user body PasswordReset true "User"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reset-password [post]
func ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	var reqBody PasswordReset
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reqBody.Password == "" || reqBody.ConfirmPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	if reqBody.Password != reqBody.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	var user model.User
	if err := config.GetDB().Where("reset_token = ?", token).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	convertedExpTime := time.Unix(user.ResetExpTime, 0)
	if convertedExpTime.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return
	}

	hash, err := argon2id.CreateHash(reqBody.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hash
	user.ResetToken = ""
	user.ResetExpTime = 0
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
