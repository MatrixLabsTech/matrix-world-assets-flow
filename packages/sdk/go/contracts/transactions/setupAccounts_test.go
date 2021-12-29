package transactions

import (
	"testing"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/sdk/go/contracts"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
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

func TestSetupAccountsTrans(t *testing.T) {
	e := contracts.NewEmulator()

	nftCode, err := contracts.GenerateNonFungibleToken()
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := contracts.GenerateLicensedNFT()
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")
	mwNFTCode, err := contracts.GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
		contracts.GetContractRoot())
	mwNFTAddr, err := e.Deploy(mwNFTCode, "MatrixWorldAssetsNFT")

	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, signer := accountKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
	require.NoError(t, err)

	s := SetupAccounts(nftAddr.String(), mwNFTAddr.String())
	tx := flow.NewTransaction().
		SetScript(s).
		SetGasLimit(9999).
		SetProposalKey(e.ServiceKey().Address, e.ServiceKey().Index, e.ServiceKey().SequenceNumber).
		SetPayer(e.ServiceKey().Address).
		AddAuthorizer(e.ServiceKey().Address)

	r, err := e.SignAndExecTrans(tx,
		[]flow.Address{e.ServiceKey().Address, addr},
		[]crypto.Signer{e.ServiceKey().Signer(), signer})
	require.NoError(t, err)
	require.False(t, r.Reverted(), r.Error)
}
