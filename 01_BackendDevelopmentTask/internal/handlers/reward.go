package handlers

import (
	"net/http"

	"github.com/demir/golang-api-backend/db"
	"github.com/gin-gonic/gin"
)

func RewardHandler(c *gin.Context) {
	walletAddress := c.Param("wallet_address")
	var amount float64

	err := db.DB.QueryRow("SELECT amount FROM staking WHERE wallet_address = ?", walletAddress).Scan(&amount)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	reward := amount * 0.05
	c.JSON(http.StatusOK, gin.H{"wallet_address": walletAddress, "rewards": reward})
}
