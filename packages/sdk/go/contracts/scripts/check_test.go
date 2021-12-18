package scripts

import (
	"github.com/MatrixLabsTech/matrix-world-assets-flow"
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/require"
)

import "testing"

func TestCheckScript(t *testing.T) {
	e := contracts.NewEmulator()

	nftCode, err := contracts.GenerateNonFungibleToken()
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := contracts.GenerateLicensedNFT()
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")
	mwNFTCode, err := contracts.GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
		getContractRoot())
	mwNFTAddr, err := e.Deploy(mwNFTCode, "MatrixWorldAssetsNFT")

	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, _ := accountKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
	require.NoError(t, err)

	s := GetCheckScript(nftAddr.String(), mwNFTAddr.String())
	r, err := e.ExecuteScript(s, [][]byte{json.MustEncode(cadence.Address(addr))})
	require.NoError(t, err)
	require.False(t, r.Reverted(), r.Error)
	t.Log(r.Value)
}
