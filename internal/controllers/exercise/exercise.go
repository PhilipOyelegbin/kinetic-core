package controller

import (
	"net/http"
	"workout_tracker/internal/config"
	model "workout_tracker/internal/model/exercise"

	"github.com/gin-gonic/gin"
)

// @Tags Exercise
// @Summary Get all exercises
// @Description Get a list of all exercises
// @Accept json
// @Produce json
// @Success 200 {array} model.Exercise
// @Failure 500 {object} map[string]string
// @Router /exercises [get]
func GetAllExercises(c *gin.Context) {
	var data []model.Exercise
	if err := config.GetDB().Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All exercises retrieved successfully", "data": data})
}

// @Tags Exercise
// @Summary Get all exercise category
// @Description Get a list of all exercise category
// @Accept json
// @Produce json
// @Success 200 {array} model.ExerciseCategory
// @Failure 500 {object} map[string]string
// @Router /exercise-categories [get]
func GetAllCategories(c *gin.Context) {
	var data []model.ExerciseCategory
	if err := config.GetDB().Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All exercise categories retrieved successfully", "data": data})
}
