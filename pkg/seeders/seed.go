package main

import (
	"log"
	"workout_tracker/internal/config"
	model "workout_tracker/internal/model/exercise"
)

var categories = []model.ExerciseCategory{
	{Name: "Strength"},
	{Name: "Cardio"},
	{Name: "Flexibility"},
}

var exercises = []model.Exercise{
	{Name: "Push-up", Description: "A classic bodyweight exercise for the chest, shoulders, and triceps.", Category: 1, MuscleGroup: "Chest"},
	{Name: "Squat", Description: "A fundamental lower body exercise that targets the quads, hamstrings, and glutes.", Category: 1, MuscleGroup: "Legs"},
	{Name: "Plank", Description: "An isometric core strength exercise that involves maintaining a position similar to a push-up for the maximum possible time.", Category: 1, MuscleGroup: "Core"},
	{Name: "Running", Description: "A popular form of cardiovascular exercise.", Category: 2, MuscleGroup: "Full Body"},
	{Name: "Hamstring Stretch", Description: "A stretch to improve flexibility in the back of the thigh.", Category: 2, MuscleGroup: "Legs"},
	{Name: "Bicep Curl", Description: "A weight training exercise that targets the biceps.", Category: 1, MuscleGroup: "Arms"},
	{Name: "Pull-up", Description: "An upper-body strength exercise where the body is pulled up.", Category: 1, MuscleGroup: "Back"},
}

func main() {
	for _, category := range categories {
		result := config.GetDB().Create(&category)
		if result.Error != nil {
			log.Printf("Could not create category '%s': %v\n", category.Name, result.Error)
		} else {
			log.Printf("Successfully seeded category: %s\n", category.Name)
		}
	}

	for _, exercise := range exercises {
		result := config.GetDB().Create(&exercise)
		if result.Error != nil {
			log.Printf("Could not create exercise '%s': %v\n", exercise.Name, result.Error)
		} else {
			log.Printf("Successfully seeded exercise: %s\n", exercise.Name)
		}
	}

	log.Println("Seeding complete!")
}
