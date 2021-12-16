// Package test performs tests on the contracts method
package test

import (
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"os"
	"testing"

	sdkContracts "github.com/MatrixLabsTech/matrix-world-assets-flow/packages/contracts/lib/go/contracts"
	sdkTemplates "github.com/onflow/flow-go-sdk/templates"
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

func deployMatrixWorldAssetsNFT(t *testing.T) (b *emulator.Blockchain, assetsNFTAddr flow.Address) {
	b = newBlockchain()

	// Should be able to deploy the NFT contract
	// as a new account with no keys.
	nftCode, _ := DownloadFile(NonFungibleTokenContractsBaseURL + NonFungibleTokenInterfaceFile)
	nftAddr, err := b.CreateAccount(nil, []sdkTemplates.Contract{
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
	licensedNFTAddr, err := b.CreateAccount(nil, []sdkTemplates.Contract{
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
	assetsNFTCode := sdkContracts.Generate("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(), getContractRoot())
	t.Logf("assetsNFTCode contract: %s", assetsNFTCode)
	assetsNFTAddr, err = b.CreateAccount(nil, []sdkTemplates.Contract{
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
	return
}

// TestNFTDeployment tests the deployment of MatrixWorld asset NFT
func TestNFTDeployment(t *testing.T) {
	deployMatrixWorldAssetsNFT(t)
}

// TestNFTSetupAccount tests the NFT setup_account
func TestNFTSetupAccount(t *testing.T) {
	b, _ := deployMatrixWorldAssetsNFT(t)
	var script []byte
	var args [][]byte
	executeScriptAndCheck(t, b, script, args)
}
