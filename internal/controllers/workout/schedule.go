package controllers

import (
	"net/http"
	"time"
	"workout_tracker/internal/config"
	model "workout_tracker/internal/model/workout"
	"workout_tracker/pkg/utils"

	"github.com/gin-gonic/gin"
)

type WorkoutSchedule struct {
	WorkoutPlanId int64     `json:"workout_plan_id"`
	ScheduledDate time.Time `json:"scheduled_date"`
	Status        string    `json:"status"`
	CompletedDate time.Time `json:"completed_date"`
}

type WorkoutScheduleDetails struct {
	ScheduledDate time.Time         `json:"scheduled_date"`
	Status        string            `json:"status"`
	CompletedDate time.Time         `json:"completed_date"`
	WorkoutPlan   model.WorkoutPlan `json:"workout_plan"`
}

// @Tags Workout
// @Summary Get user workout schedule
// @Description Get the workout schedule of the authenticated user
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutSchedule
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/schedules [get]
func GetMyWorkoutSchedules(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var schedules []WorkoutSchedule
	if err := config.GetDB().Find(&schedules, map[string]interface{}{"user_id": userId}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schedules"})
		return
	}
	if len(schedules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No workout schedules found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All schedules retrieved successfully", "data": schedules})
}

// @Tags Workout
// @Summary Get user workout schedule by id
// @Description Get the workout schedule of the authenticated user by id
// @Param id path int true "Workout ID"
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutSchedule
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/schedules/{id} [get]
func GetScheduleByID(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	scheduleId := c.Params.ByName("id")
	if scheduleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Schedule Id is required!"})
		return
	}

	var schedule model.WorkoutSchedule
	if err := config.GetDB().First(&schedule, map[string]interface{}{"ID": scheduleId, "user_id": userId}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout schedule not found"})
		return
	}

	var workout model.WorkoutPlan
	if err := config.GetDB().First(&workout, map[string]interface{}{"ID": schedule.WorkoutPlanId}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout plan not found"})
		return
	}

	response := WorkoutScheduleDetails{
		ScheduledDate: schedule.ScheduledDate,
		Status:        schedule.Status,
		CompletedDate: schedule.CompletedDate,
		WorkoutPlan:   workout,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedules retrieved successfully", "data": response})
}

// @Tags Workout
// @Summary Filter user workout schedule by status
// @Description Filter the workout schedule of the authenticated user by status
// @Param status query string true "Workout Status"
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutSchedule
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/schedules/status [get]
func FilterByStatus(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required!"})
		return
	}

	var schedules []WorkoutSchedule
	if err := config.GetDB().Find(&schedules, map[string]interface{}{"user_id": userId, "status": status}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schedules"})
		return
	}

	if len(schedules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No workout schedules found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedules retrieved successfully", "data": schedules})
}

// @Tags Workout
// @Summary Create new user workout schedule
// @Description Create a new workout schedule for the authenticated user
// @Param schedule body WorkoutSchedule true "Workout Schedule"
// @Accept json
// @Produce json
// @Success 200 {object} WorkoutSchedule
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/schedules [post]
func CreateSchedule(c *gin.Context) {
	userId, err := utils.ExtractUserIdFromJWTToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	schedule := model.WorkoutSchedule{}
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	schedule.UserId = userId
	if err := config.GetDB().Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created successfully", "data": schedule})
}
