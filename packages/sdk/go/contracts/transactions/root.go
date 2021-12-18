package transactions

import (
	"os"
	"path/filepath"
)

const (
	defaultTransactionsRoot = "../../../../contracts/cadence/transactions/"
	defaultContractRoot     = "../../../../contracts/cadence/contracts/"
)

var (
	transRoot string
)

// GetRoot reads from environment or use default value if not set
func GetRoot() string {
	if transRoot == "" {
		transRoot = os.Getenv("TRANSACTION_ROOT")
		if transRoot == "" {
			transRoot = defaultTransactionsRoot
		}
	}
	if transRoot == "" {
		panic("root is not valid")
	}
	return filepath.Clean(transRoot)
}

func SetRoot(root string) {
	if root == "" {
		panic("root is not valid")
	}
	transRoot = filepath.Clean(root)
}

// get CONTRACT_ROOT from environment or use default value if not set
func getContractRoot() string {
	contractRoot := os.Getenv("CONTRACT_ROOT")
	if contractRoot == "" {
		contractRoot = defaultContractRoot
	}
	return contractRoot
}
