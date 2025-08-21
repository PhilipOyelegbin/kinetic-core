package main

import (
	"log"
	"net/http"
	"os"
	routes "workout_tracker/api"
	"workout_tracker/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "workout_tracker/docs"
)

// @title Kinetic Core API
// @version 1.0
// @description This project involves creating a backend system for a workout tracker application where users can sign up, log in, create workout plans, and track their progress. The system will feature JWT authentication, CRUD operations for workouts, and generate reports on past workouts.
// @termsOfService http://swagger.io/terms/
// @contact.name Philip Oyelegbin
// @contact.url https://philipoyelegbin.com.ng
// @contact.email info@philipoyelegbin.com.ng
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @Security BearerAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" and then your JWT token to authorize

func main() {
	env := config.LoadEnv()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	app := gin.Default()
	api := app.Group("/api/v1")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routes.RegisterRoutes(api)
	}


	app.GET("/", gin.HandlerFunc(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/swagger/index.html")
	}))

	err := app.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
