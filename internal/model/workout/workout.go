package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type WorkoutPlan struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	UserId      int64   `json:"user_id" gorm:"foreignKey:UserId"`
	ExerciseId  int64   `json:"exercise_id" gorm:"foreignKey:ExerciseId"`
	Sets        int64   `json:"sets" gorm:"not null"`
	Repetitions int64   `json:"repetitions" gorm:"not null"`
	Weight      float32 `json:"weight" gorm:"not null"`
	Order       int64   `json:"order" gorm:"not null"`
}

type WorkoutSchedule struct {
	gorm.Model
	UserId        int64     `json:"user_id" gorm:"foreignKey:UserId"`
	WorkoutPlanId int64     `json:"workout_plan_id" gorm:"foreignKey:WorkoutPlanId"`
	ScheduledDate time.Time `json:"scheduled_date" gorm:"not null"`
	Status        string    `json:"status" gorm:"default:'scheduled'"`
	CompletedDate time.Time `json:"completed_date"`
}
