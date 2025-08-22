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

	var schedules []model.WorkoutSchedule
	if err := config.GetDB().Find(&schedules, map[string]interface{}{"user_id": userId}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schedules"})
		return
	}

	if len(schedules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No workout schedules found"})
		return
	}
	c.JSON(http.StatusOK, schedules)
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
	c.JSON(http.StatusOK, schedule)
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

	c.JSON(http.StatusCreated, schedule)
}
