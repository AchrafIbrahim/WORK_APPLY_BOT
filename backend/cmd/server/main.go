package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/config"
	"github.com/AchrafIbrahim/WORK_APPLY_BOT/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connection successful")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
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

	r.Run(":8081")
}