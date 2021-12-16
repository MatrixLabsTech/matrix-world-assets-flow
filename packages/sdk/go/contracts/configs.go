package contracts

import "os"

const (
	NonFungibleTokenContractsBaseURL = "https://raw.githubusercontent.com/onflow/flow-nft/master/contracts/"
	NonFungibleTokenInterfaceFile    = "NonFungibleToken.cdc"
	LicensedNFTInterfaceURL = "https://raw.githubusercontent.com/rarible/flow-contracts/main/contracts/LicensedNFT.cdc"

	defaultContractRoot = "../../../contracts/cadence/contracts/"
	emulatorFTAddress        = "ee82856bf20e2aa6"
	emulatorFlowTokenAddress = "0ae53cb6e3f42a79"
)

var (
	// get CONTRACT_ROOT from environment or use default value if not set
	contractRoot = getContractRoot()
)

// get CONTRACT_ROOT from environment or use default value if not set
func getContractRoot() string {
	contractRoot := os.Getenv("CONTRACT_ROOT")
	if contractRoot == "" {
		contractRoot = defaultContractRoot
	}
	return contractRoot
}