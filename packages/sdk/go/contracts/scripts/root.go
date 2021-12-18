package scripts

import (
	"os"
	"path/filepath"
)

const (
	defaultScriptRoot   = "../../../../contracts/cadence/scripts/"
	defaultContractRoot = "../../../../contracts/cadence/contracts/"
)

var (
	scriptRoot string
)

//	GetRoot returns folders of the script folder,
//	or uses SCRIPT_ROOT from environment if not sets,
//	or uses default if env is empty
func GetRoot() string {
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

func SetRoot(root string) {
	if root == "" {
		panic("root is not valid")
	}
	scriptRoot = filepath.Clean(root)
}

// get CONTRACT_ROOT from environment or use default value if not set
func getContractRoot() string {
	contractRoot := os.Getenv("CONTRACT_ROOT")
	if contractRoot == "" {
		contractRoot = defaultContractRoot
	}
	return contractRoot
}
