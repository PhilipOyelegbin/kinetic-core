package model

import "github.com/jinzhu/gorm"

type ExerciseCategory struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}

type Exercise struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Category    int    `json:"category"`
	MuscleGroup string `json:"muscle_group" gorm:"index"`
}
