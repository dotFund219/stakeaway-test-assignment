package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/demir/golang-sepolia-staking/stake"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get values from environment variables
	rpcURL := os.Getenv("INFURA_PROJECT_ID")
	privateKey := os.Getenv("PRIVATE_KEY")
	stakingContract := os.Getenv("STAKING_CONTRACT")

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	accountAddress := getAddressFromPrivateKey(privateKey)
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(accountAddress), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	fmt.Printf("Account Balance: %s ETH\n", weiToEther(balance).String())

	txHash, err := sendTransaction(client, privateKey, stakingContract, big.NewInt(100000000000000)) // 0.0001 ETH
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Printf("Transaction sent! Hash: %s\n", txHash)
}

func getAddressFromPrivateKey(privKeyHex string) string {
	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKey).Hex()
}

func sendTransaction(client *ethclient.Client, privKeyHex, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	// Load the smart contract
	contractAddressCommon := common.HexToAddress(toAddress)
	instance, err := stake.NewStake(contractAddressCommon, client)
	if err != nil {
		log.Fatalf("Failed to load smart contract: %v", err)
	}

	// Get the next available nonce for the account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	// Get the current gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Set up the transaction options
	chainID := big.NewInt(11155111) // Sepolia Testnet Chain ID (Replace with your network's Chain ID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = amount            // Staking 0.0001 ETH
	auth.GasLimit = uint64(300000) // Set appropriate gas limit
	auth.GasPrice = gasPrice

	// Call stake function
	tx, err := instance.Stake(auth)
	if err != nil {
		log.Fatalf("Failed to send stake transaction: %v", err)
	}

	fmt.Println("Transaction sent! Tx Hash:", tx.Hash().Hex())

	// Wait for confirmation
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Transaction mining failed: %v", err)
	}

	fmt.Println("Transaction confirmed with status:", receipt.Status)
	return tx.Hash().Hex(), nil
}

func weiToEther(wei *big.Int) *big.Float {
	ether := new(big.Float).SetInt(wei)
	return new(big.Float).Quo(ether, big.NewFloat(1e18))
}
