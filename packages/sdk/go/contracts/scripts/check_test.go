package scripts

import (
	"github.com/MatrixLabsTech/matrix-world-assets-flow"
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

import "testing"

func TestCheckScript(t *testing.T) {
	e := contracts.NewEmulator()

	nftCode, err := contracts.GenerateNonFungibleToken()
	assert.NoError(t, err)
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := contracts.GenerateLicensedNFT()
	assert.NoError(t, err)
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")

	mwNFTCode, err := contracts.GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
		getContractRoot())
	assert.NoError(t, err)
	mwAssetNFTAddr, err := e.Deploy(mwNFTCode, "MatrixWorldAssetsNFT")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	t.Log("mwAssetNFTAddr", mwAssetNFTAddr)
	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, _ := accountKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
	require.NoError(t, err)

	s := CheckScript(nftAddr.String(), mwAssetNFTAddr.String())
	r, err := e.ExecuteScript(s, [][]byte{json.MustEncode(cadence.Address(addr))})
	if r.Reverted() {
		t.Fatal(err)
	}
}
