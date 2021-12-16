package scripts

import "os"

const (
	defaultScriptRoot = "../../../../contracts/cadence/scripts/"
	defaultContractRoot = "../../../../contracts/cadence/contracts/"
)

// get SCRIPT_ROOT from environment or use default value if not set
func getScriptRoot() string {
	contractRoot := os.Getenv("SCRIPT_ROOT")
	if contractRoot == "" {
		contractRoot = defaultScriptRoot
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