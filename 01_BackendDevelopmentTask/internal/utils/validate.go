package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

// IsValidEthereumAddress checks if a given string is a valid Ethereum address
func IsValidEthereumAddress(address string) bool {
	return common.IsHexAddress(address)
}
