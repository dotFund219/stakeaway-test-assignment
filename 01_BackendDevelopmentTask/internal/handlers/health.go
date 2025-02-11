package handlers

import (
	"net/http"
	"time"

	"github.com/demir/golang-api-backend/config"
	"github.com/demir/golang-api-backend/db"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	if err := db.DB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Database connection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Service is healthy", "uptime": time.Since(config.StartTime).String()})
}
