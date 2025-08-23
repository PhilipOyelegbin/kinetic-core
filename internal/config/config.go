package config

import (
	"log"
	"os"
	exercise "workout_tracker/internal/model/exercise"
	user "workout_tracker/internal/model/user"
	workout "workout_tracker/internal/model/workout"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func LoadEnv() error {
	return godotenv.Load()
}

func connectDatabase() {
	LoadEnv()
	conn, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = conn
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return DB
}

func init() {
	connectDatabase()
	DB = GetDB()
	DB.AutoMigrate(&exercise.ExerciseCategory{})
	DB.AutoMigrate(&exercise.Exercise{})
	DB.AutoMigrate(&user.User{})
	DB.AutoMigrate(&workout.WorkoutPlan{})
	DB.AutoMigrate(&workout.WorkoutSchedule{})
	log.Println("Database migrated and connected successfully")
}
