package transactions

import "os"

const (
	defaultTransactionsRoot = "../../../../contracts/cadence/transactions/"
	defaultContractRoot = "../../../../contracts/cadence/contracts/"
)

// get TRANSACTION_ROOT from environment or use default value if not set
func getTransactionsRoot() string {
	contractRoot := os.Getenv("TRANSACTION_ROOT")
	if contractRoot == "" {
		contractRoot = defaultTransactionsRoot
	}
	return contractRoot
}

// get CONTRACT_ROOT from environment or use default value if not set
func getContractRoot() string {
	contractRoot := os.Getenv("CONTRACT_ROOT")
	if contractRoot == "" {
		contractRoot = defaultContractRoot
	}
	return contractRoot
}