package handlers

import (
	"net/http"
	"sync"

	"github.com/demir/golang-api-backend/db"
	"github.com/demir/golang-api-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

var mu sync.Mutex

func StakeHandler(c *gin.Context) {
	var request struct {
		WalletAddress string  `json:"wallet_address"`
		Amount        float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be positive"})
		return
	}

	if !utils.IsValidEthereumAddress(request.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet address is not valid"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	_, err := db.DB.Exec("INSERT INTO staking (wallet_address, amount) VALUES (?, ?) ON CONFLICT(wallet_address) DO UPDATE SET amount = amount + excluded.amount", request.WalletAddress, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staking successful"})
}
