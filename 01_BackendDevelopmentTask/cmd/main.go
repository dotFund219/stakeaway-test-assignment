package main

import (
	"log"
	"time"

	"github.com/demir/golang-api-backend/config"
	"github.com/demir/golang-api-backend/db"
	"github.com/demir/golang-api-backend/internal/handlers"
	"github.com/demir/golang-api-backend/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.StartTime = time.Now()
	db.InitDB()

	router := gin.Default()

	router.Use(middlewares.PrometheusMiddleware())

	// Define API routes
	router.GET("/ping", handlers.PingHandler)
	router.POST("/stake", handlers.StakeHandler)
	router.GET("/rewards/:wallet_address", handlers.RewardHandler)
	router.GET("/health", handlers.HealthHandler)
	router.POST("/validators", handlers.CreateValidatorRequest)
	router.GET("/validators/:id", handlers.GetRequestStatus)

	// Expose Prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
