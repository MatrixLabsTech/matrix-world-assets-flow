package test

import (
	"os"
	"testing"

	sdkcontracts "github.com/MatrixLabsTech/matrix-world-assets-flow/packages/contracts/lib/go/contracts"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/stretchr/testify/assert"
)

const (
	NonFungibleTokenContractsBaseURL = "https://raw.githubusercontent.com/onflow/flow-nft/master/contracts/"
	NonFungibleTokenInterfaceFile    = "NonFungibleToken.cdc"

	LicensedNFTInterfaceURL = "https://raw.githubusercontent.com/rarible/flow-contracts/main/contracts/LicensedNFT.cdc"

	emulatorFTAddress        = "ee82856bf20e2aa6"
	emulatorFlowTokenAddress = "0ae53cb6e3f42a79"
	// get CONTRACT_ROOT from environment
)

// get CONTRACT_ROOT from environment or use default value if not set
func getContractRoot() string {
	contractRoot := os.Getenv("CONTRACT_ROOT")
	if contractRoot == "" {
		contractRoot = "../../../cadence/contracts/"
	}
	return contractRoot
}

var (
	// get CONTRACT_ROOT from environment or use default value if not set
	contractRoot = getContractRoot()
)

// This test is for testing the deployment the topshot smart contracts
func TestNFTDeployment(t *testing.T) {
	b := newBlockchain()

	// Should be able to deploy the NFT contract
	// as a new account with no keys.
	nftCode, _ := DownloadFile(NonFungibleTokenContractsBaseURL + NonFungibleTokenInterfaceFile)
	nftAddr, err := b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "NonFungibleToken",
			Source: string(nftCode),
		},
	})
	t.Logf("NFT address: %s", nftAddr)
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// deploy LicensedNFTInterface
	licensedNFTCode, _ := DownloadFile(LicensedNFTInterfaceURL)
	licensedNFTAddr, err := b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "LicensedNFT",
			Source: string(licensedNFTCode),
		},
	})
	t.Logf("LicensedNFT address: %s", licensedNFTAddr)
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy AssetNFT contract
	// as a new account with no keys.
	assetsNFTCode := sdkcontracts.GenerateTopShotContract("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(), getContractRoot())
	t.Logf("assetsNFTCode contract: %s", assetsNFTCode)
	assetsNFTAddr, err := b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "MatrixWorldAssetsNFT",
			Source: string(assetsNFTCode),
		},
	})
	t.Logf("MatrixWorldAssetsNFT address: %s", assetsNFTAddr)
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)

}
