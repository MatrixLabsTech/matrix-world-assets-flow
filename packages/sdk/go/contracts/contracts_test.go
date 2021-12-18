// Package test performs tests on the contracts method
package contracts

import (
	"github.com/MatrixLabsTech/matrix-world-assets-flow/scripts"
	"github.com/MatrixLabsTech/matrix-world-assets-flow/transactions"
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	scriptsRoot      = "../../../contracts/cadence/scripts/"
	transactionsRoot = "../../../contracts/cadence/transactions/"
)

func init() {
	scripts.SetRoot(scriptsRoot)
	transactions.SetRoot(transactionsRoot)
}

func TestNonFungibleToken(t *testing.T) {
	e := NewEmulator()

	code, err := GenerateNonFungibleToken()
	require.NoError(t, err)
	addr, err := e.Deploy(code, "NonFungibleToken")
	require.NoError(t, err)

	t.Logf("NonFungibleToken address: %s", addr)
}

func TestLicensedNFT(t *testing.T) {
	e := NewEmulator()

	code, err := GenerateLicensedNFT()
	require.NoError(t, err)
	addr, err := e.Deploy(code, "LicensedNFT")
	require.NoError(t, err)
	t.Logf("LicensedNFT address: %s", addr)
}

func TestMatrixWorldAssetsNFT(t *testing.T) {
	e := NewEmulator()

	nftCode, err := GenerateNonFungibleToken()
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := GenerateLicensedNFT()
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")

	code, err := GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(), getContractRoot())
	require.NoError(t, err)
	mwNFTAddr, err := e.Deploy(code, "MatrixWorldAssetsNFT")
	require.NoError(t, err)
	t.Logf("MatrixWorldAssetsNFT address: %s", mwNFTAddr)
}

func TestExecScriptsTransactions(t *testing.T) {
	e := NewEmulator()

	nftCode, err := GenerateNonFungibleToken()
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := GenerateLicensedNFT()
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")
	mwNFTCode, err := GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
		getContractRoot())
	mwNFTAddr, err := e.Deploy(mwNFTCode, "MatrixWorldAssetsNFT")

	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, signer := accountKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
	require.NoError(t, err)

	ss := scripts.GetCheckScript(nftAddr.String(), mwNFTAddr.String())
	sr, err := e.ExecuteScript(ss, [][]byte{json.MustEncode(cadence.Address(mwNFTAddr))})
	require.NoError(t, err)
	require.False(t, sr.Reverted(), sr.Error)

	ts := transactions.SetupAccounts(nftAddr.String(), mwNFTAddr.String())
	tx := e.CreateTrans(ts, e.ServiceKey().Address)

	tr, err := e.SignAndExecTrans(tx,
		[]flow.Address{e.ServiceKey().Address, addr},
		[]crypto.Signer{e.ServiceKey().Signer(), signer})
	require.NoError(t, err)
	require.False(t, tr.Reverted(), tr.Error)

	ss = scripts.GetCheckScript(nftAddr.String(), mwNFTAddr.String())
	sr, err = e.ExecuteScript(ss, [][]byte{json.MustEncode(cadence.Address(mwNFTAddr))})
	require.NoError(t, err)
	require.False(t, sr.Reverted(), sr.Error)
	require.EqualValues(t, true, sr.Value)
}
