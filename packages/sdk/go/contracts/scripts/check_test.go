package scripts

import (
	"fmt"
	"testing"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/require"
)

const (
	scriptsRoot      = "../../../../contracts/cadence/scripts/"
	transactionsRoot = "../../../../contracts/cadence/transactions/"
	contractRoot     = "../../../../contracts/cadence/contracts/"
)

func init() {
	contracts.SetScriptRoot(scriptsRoot)
	contracts.SetTransRoot(transactionsRoot)
	contracts.SetContractRoot(contractRoot)
}

func TestCheckScript(t *testing.T) {
	e := contracts.NewEmulator()

	nftCode, err := contracts.GenerateNonFungibleToken()
	require.NoError(t, err)

	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	require.NoError(t, err)

	_, err = e.CommitBlock()
	require.NoError(t, err)

	licensedNFTCode, err := contracts.GenerateLicensedNFT()
	require.NoError(t, err)

	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")
	require.NoError(t, err)

	_, err = e.CommitBlock()
	require.NoError(t, err)

	mwNFTCode, err := contracts.GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
		contracts.GetContractRoot())
	require.NoError(t, err)

	mwNFTAddr, err := e.Deploy(mwNFTCode, "MatrixWorldAssetsNFT")
	require.NoError(t, err)

	_, err = e.CommitBlock()
	require.NoError(t, err)

	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, _ := accountKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
	require.NoError(t, err)

	s := GetCheckScript(nftAddr.String(), mwNFTAddr.String())
	// bytes to string
	fmt.Print(string(s))
	r, err := e.ExecuteScript(s, [][]byte{json.MustEncode(cadence.Address(addr))})
	require.NoError(t, err)
	require.False(t, r.Reverted(), r.Error)

	_, err = e.CommitBlock()
	require.NoError(t, err)

	t.Log(r.Value)
}
