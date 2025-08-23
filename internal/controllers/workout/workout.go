package controllers

import (
	"log"
	"net/http"
	"workout_tracker/internal/config"
	exeModel "workout_tracker/internal/model/exercise"
	model "workout_tracker/internal/model/workout"
	"workout_tracker/pkg/utils"

	"github.com/gin-gonic/gin"
)

type WorkoutPlan struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ExerciseId  int64   `json:"exercise_id"`
	Sets        int64   `json:"sets"`
	Repetitions int64   `json:"repetitions"`
	Weight      float32 `json:"weight"`
	Order       int64   `json:"order"`
}
type WorkoutPlanDetails struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ExerciseId  exeModel.Exercise `json:"exercise"`
	Sets        int64             `json:"sets"`
	Repetitions int64             `json:"repetitions"`
	Weight      float32           `json:"weight"`
	Order       int64             `json:"order"`
}
type WorkoutReport struct {
	WorkoutName   string  `json:"workout_name"`
	TotalReps     uint    `json:"total_reps"`
	AvgWeight     float64 `json:"average_weight"`
	TotalWorkouts int64   `json:"total_workouts"`
}

// @Tags Workout
// @Summary Get user workout plan
// @Description Get the workout plan of the authenticated user
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts [get]
func GetMyWorkouts(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var workouts []model.WorkoutPlan
	if err := config.GetDB().Find(&workouts, map[string]interface{}{"user_id": userId}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workouts"})
		return
	}
	if len(workouts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No workouts found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All workouts retrieved successfully", "data": workouts})
}

// @Tags Workout
// @Summary Get user workout plan by id
// @Description Get the workout plan of the authenticated user by id
// @Param id path int true "Workout ID"
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/{id} [get]
func GetWorkoutByID(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	workoutId := c.Params.ByName("id")
	if workoutId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workout Id is required!"})
		return
	}

	var workout model.WorkoutPlan
	if err := config.GetDB().First(&workout, map[string]interface{}{"ID": workoutId, "user_id": userId}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout plan not found"})
		return
	}

	var exercise exeModel.Exercise
	if err := config.GetDB().First(&exercise, workout.ExerciseId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	response := WorkoutPlanDetails{
		Name:        workout.Name,
		Description: workout.Description,
		ExerciseId:  exercise,
		Sets:        workout.Sets,
		Repetitions: workout.Repetitions,
		Weight:      workout.Weight,
		Order:       workout.Order,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workout retrieved successfully", "data": response})
}

// @Tags Workout
// @Summary Create user workout plan
// @Description Create a new workout plan for the authenticated user
// @Param workout body WorkoutPlan true "Workout"
// @Accept json
// @Produce json
// @Success 201 {object} WorkoutPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts [post]
func CreateWorkout(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var reqBody model.WorkoutPlan
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	reqBody.UserId = userId
	if err := config.GetDB().Create(&reqBody).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create workout"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Workout created successfully", "data": reqBody})
}

// @Tags Workout
// @Summary Update user workout plan
// @Description Update a workout plan for the authenticated user
// @Param id path int true "Workout ID"
// @Param workout body WorkoutPlan true "Workout"
// @Accept json
// @Produce json
// @Success 202 {object} WorkoutPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/{id} [patch]
func UpdateWorkout(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	workoutId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workout Id is required!"})
		return
	}

	var reqBody map[string]interface{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	delete(reqBody, "user_id")
	delete(reqBody, "id")
	if len(reqBody) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return
	}

	result := config.GetDB().Model(&model.WorkoutPlan{}).Where(map[string]interface{}{"ID": workoutId, "user_id": userId}).Updates(reqBody)
	if result.Error != nil {
		log.Printf("Database error updating workout: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update workout plan"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout plan not found or not authorized"})
		return
	}

	var updatedWorkout model.WorkoutPlan
	if err := config.GetDB().First(&updatedWorkout, workoutId).Error; err != nil {
		log.Printf("Error fetching updated workout: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated workout plan"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Workout plan updated successfully", "data": updatedWorkout})
}

// @Tags Workout
// @Summary Delete user workout plan by id
// @Description Delete the workout plan of the authenticated user by id
// @Param id path int true "Workout ID"
// @Accept json
// @Produce json
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/{id} [delete]
func DeleteWorkout(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	workoutId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workout Id is required!"})
		return
	}

	result := config.GetDB().Delete(&model.WorkoutPlan{}, map[string]interface{}{"ID": workoutId, "user_id": userId})
	if result.Error != nil {
		log.Printf("Database error deleting workout: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete workout plan"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout plan not found or not authorized"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Workout plan deleted"})
}

// @Tags Workout
// @Summary Get user workout reports
// @Description Get the workout reports of the authenticated user
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutReport
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/reports [get]
func GenerateWorkoutReport(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var report []WorkoutReport
	selectStatement := "name as workout_name, SUM(repetitions) as total_reps, AVG(weight) as avg_weight, COUNT(*) as total_workouts"
	result := config.GetDB().Model(&WorkoutPlan{}).Select(selectStatement).Where("user_id = ?", userId).Group("name").Scan(&report)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}
	if len(report) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No workout data found for this user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workout report generated successfully", "data": report})
}
