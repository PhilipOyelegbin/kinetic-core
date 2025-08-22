package routes

import (
	auth "workout_tracker/internal/controllers/auth"
	exercise "workout_tracker/internal/controllers/exercise"
	user "workout_tracker/internal/controllers/user"
	workout "workout_tracker/internal/controllers/workout"

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
	api.GET("/users", user.GetMyProfile)
	api.PATCH("/users/change-password", user.UpdatePassword)

	// exercise routes
	api.GET("/exercises", exercise.GetAllExercises)
	api.GET("/exercise-categories", exercise.GetAllCategories)

	// workout routes
	api.GET("/workouts", workout.GetMyWorkouts)
	api.POST("/workouts", workout.CreateWorkout)
	api.GET("/workouts/:id", workout.GetWorkoutByID)
	api.PATCH("/workouts/:id", workout.UpdateWorkout)
	api.DELETE("/workouts/:id", workout.DeleteWorkout)
	api.GET("/workouts/schedules", workout.GetMyWorkoutSchedules)
	api.POST("/workouts/schedules", workout.CreateSchedule)
	api.GET("/workouts/schedules/:id", workout.GetScheduleByID)
}
