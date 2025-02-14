package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/demir/golang-api-backend/db"
	"github.com/demir/golang-api-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

var mutex sync.Mutex

func CreateValidatorRequest(c *gin.Context) {
	var request struct {
		NumValidators int    `json:"num_validators"`
		FeeRecipient  string `json:"fee_recipient"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if request.NumValidators <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The number of validators must be positive"})
		return
	}

	if !utils.IsValidEthereumAddress(request.FeeRecipient) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fee recipient address is not valid format"})
		return
	}

	requestID := uuid.New().String()

	_, err := db.DB.Exec("INSERT INTO validator_requests (id, num_validators, fee_recipient, status) VALUES (?, ?, ?, ?)", requestID, request.NumValidators, request.FeeRecipient, "started")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	go processValidatorRequest(requestID, request.NumValidators)

	c.JSON(http.StatusOK, gin.H{
		"request_id": requestID,
		"message":    "Validator creation in progress",
	})
}

func GetRequestStatus(c *gin.Context) {
	requestID := c.Param("id")
	var status string
	var keys string
	err := db.DB.QueryRow("SELECT status, keys FROM validator_requests WHERE id = ?", requestID).Scan(&status, &keys)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	response := gin.H{"status": status}
	if status == "successful" {
		response["keys"] = keys
	}
	if status == "failed" {
		response["message"] = "Error processing request"
	}
	c.JSON(http.StatusOK, response)
}

func processValidatorRequest(requestID string, numValidators int) {
	time.Sleep(time.Millisecond * 100) // Simulate processing time
	mutex.Lock()
	defer mutex.Unlock()

	keys := generateKeys(numValidators)
	keysJSON, _ := json.Marshal(keys)
	_, err := db.DB.Exec("UPDATE validator_requests SET status = ?, keys = ? WHERE id = ?", "successful", string(keysJSON), requestID)
	if err != nil {
		_, _ = db.DB.Exec("UPDATE validator_requests SET status = ? WHERE id = ?", "failed", requestID)
	}
}

func generateKeys(count int) []string {
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 20)
		keys[i] = fmt.Sprintf("key-%d-%d", i, rand.Intn(100000))
	}
	return keys
}
