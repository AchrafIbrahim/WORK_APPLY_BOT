package main

import (
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/config"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/database"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/handlers"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/repositories"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connection successful")

	applicationRepository := repositories.NewApplicationRepository(db)
	applicationService := services.NewApplicationService(
		applicationRepository,
	)
	applicationHandler := handlers.NewApplicationHandler(
		applicationService,
	)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "WORK APPLY BOT is running",
		})
	})

	r.GET("/health/db", func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "error",
				"message": "Database connection failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Database connection is healthy",
		})
	})
	r.POST("/applications", applicationHandler.CreateApplication)

	r.Run(":8081")
}
