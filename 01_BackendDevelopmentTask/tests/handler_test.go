package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/demir/golang-api-backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a Gin router with routes
func setupRouter() *gin.Engine {
	// db.InitDB()
	router := gin.Default()
	router.GET("/ping", handlers.PingHandler)
	router.POST("/stake", handlers.StakeHandler)
	router.GET("/rewards/:wallet_address", handlers.RewardHandler)
	router.GET("/health", handlers.HealthHandler)
	return router
}

// Test /ping endpoint
func TestPingHandler(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())
}

// Test /stake endpoint - Valid Request
func TestStakeHandlerValid(t *testing.T) {
	router := setupRouter()

	requestBody, _ := json.Marshal(map[string]interface{}{
		"wallet_address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
		"amount":         10,
	})

	req, _ := http.NewRequest("POST", "/stake", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Staking successful"}`, w.Body.String())
}

// Test /stake endpoint - Invalid Request (Negative Amount)
func TestStakeHandlerInvalidAmount(t *testing.T) {
	router := setupRouter()

	requestBody, _ := json.Marshal(map[string]interface{}{
		"wallet_address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
		"amount":         -5.0,
	})

	req, _ := http.NewRequest("POST", "/stake", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Amount must be positive"}`, w.Body.String())
}

// Test /stake endpoint - Invalid Request (Negative Amount)
func TestStakeHandlerInvalidWalletAddress(t *testing.T) {
	router := setupRouter()

	requestBody, _ := json.Marshal(map[string]interface{}{
		"wallet_address": "0xaaa",
		"amount":         5.0,
	})

	req, _ := http.NewRequest("POST", "/stake", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Wallet address is not valid"}`, w.Body.String())
}

// Test /stake endpoint - Invalid JSON Body
func TestStakeHandlerInvalidJSON(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/stake", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Invalid input"}`, w.Body.String())
}
