package contracts

import (
	"os"
	"path/filepath"
)

const (
	defaultTransactionsRoot = "../../../contracts/cadence/transactions/"
	defaultContractRoot     = "../../../contracts/cadence/contracts/"
	defaultScriptRoot       = "../../../contracts/cadence/scripts/"
)

var (
	transRoot    string
	scriptRoot   string
	contractRoot string
)

// GetRoot reads from environment or use default value if not set
func GetTransRoot() string {
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

func SetTransRoot(root string) {
	if root == "" {
		panic("root is not valid")
	}
	transRoot = filepath.Clean(root)
}

func GetScriptRoot() string {
	if scriptRoot == "" {
		scriptRoot = os.Getenv("SCRIPT_ROOT")
		if scriptRoot == "" {
			scriptRoot = defaultScriptRoot
		}
	}
	if scriptRoot == "" {
		panic("root is not valid")
	}
	return filepath.Clean(scriptRoot)
}

func SetScriptRoot(root string) {
	if root == "" {
		panic("root is not valid")
	}
	scriptRoot = filepath.Clean(root)
}

// get CONTRACT_ROOT from environment or use default value if not set
func GetContractRoot() string {
	if contractRoot == "" {
		contractRoot = os.Getenv("CONTRACT_ROOT")
		if contractRoot == "" {
			contractRoot = defaultContractRoot
		}
	}
	if contractRoot == "" {
		panic("root is not valid")
	}
	return contractRoot
}

func SetContractRoot(root string) {
	if root == "" {
		panic("root is not valid")
	}
	contractRoot = filepath.Clean(root)
}
