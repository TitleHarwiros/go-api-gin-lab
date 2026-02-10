package main

import (
	"go-api-gin/handlers"
	"go-api-gin/models"
	"go-api-gin/repositories"
	"go-api-gin/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Database
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate
	db.AutoMigrate(&models.Student{})

	// Setup Layers (Dependency Injection)
	repo := repositories.NewStudentRepository(db)
	svc := services.NewStudentService(repo)
	handler := handlers.NewStudentHandler(svc)

	r := gin.Default()

	// Routes
	r.GET("/students", handler.GetStudents)
	r.POST("/students", handler.PostStudent)
	r.PUT("/students/:id", handler.UpdateStudent)    // Challenge 1
	r.DELETE("/students/:id", handler.DeleteStudent) // Challenge 2

	r.Run(":8080")
}
