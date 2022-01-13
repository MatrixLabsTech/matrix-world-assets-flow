package mint

import (
	"testing"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts/lib"
	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts/transactions"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/require"
)

const (
	scriptsRoot      = "../../../../../contracts/cadence/scripts/"
	transactionsRoot = "../../../../../contracts/cadence/transactions/"
	contractRoot     = "../../../../../contracts/cadence/contracts/"
)

func init() {
	contracts.SetScriptRoot(scriptsRoot)
	contracts.SetTransRoot(transactionsRoot)
	contracts.SetContractRoot(contractRoot)
}

func TestMatrixWorldAssetsNFT(t *testing.T) {
	e := contracts.NewEmulator()

	nftCode, err := contracts.GenerateNonFungibleToken()
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := contracts.GenerateLicensedNFT()
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")

	code, err := contracts.GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(), contracts.GetContractRoot())
	require.NoError(t, err)

	// Create a new user account
	accountKeys := test.AccountKeyGenerator()
	pk, signer := accountKeys.NewWithSigner()

	assetsAddr, err := e.CreateAccount([]*flow.AccountKey{pk}, []sdktemplates.Contract{
		{
			Name:   "MatrixWorldAssetsNFT",
			Source: string(code),
		},
	})

	_, err = e.CommitBlock()
	require.NoError(t, err)

	t.Logf("MatrixWorldAssetsNFT address: %s", assetsAddr)

	// SetupAccount
	// Create a new user account
	userKeys := test.AccountKeyGenerator()
	userPk, userSigner := userKeys.NewWithSigner()
	addr, err := e.CreateAccount([]*flow.AccountKey{userPk}, nil)
	require.NoError(t, err)

	s := transactions.SetupAccounts(nftAddr.String(), assetsAddr.String())
	tx := flow.NewTransaction().
		SetScript(s).
		SetGasLimit(9999).
		SetProposalKey(e.ServiceKey().Address, e.ServiceKey().Index, e.ServiceKey().SequenceNumber).
		SetPayer(e.ServiceKey().Address).
		AddAuthorizer(addr)

	r, err := e.SignAndExecTrans(tx,
		[]flow.Address{e.ServiceKey().Address, addr},
		[]crypto.Signer{e.ServiceKey().Signer(), userSigner})

	require.NoError(t, err)
	require.False(t, r.Reverted(), r.Error)

	_, err = e.CommitBlock()
	require.NoError(t, err)
	t.Logf("SetupAccounts: %s", r.Events)

	// Mint
	mintCode := GenerateMintNFTTransaction(nftAddr.String(), assetsAddr.String())

	mintTx := flow.NewTransaction().
		SetScript(mintCode).
		SetGasLimit(9999).
		SetProposalKey(e.ServiceKey().Address, e.ServiceKey().Index, e.ServiceKey().SequenceNumber).
		SetPayer(e.ServiceKey().Address).
		AddAuthorizer(assetsAddr)

	recipientArray := make([]cadence.Value, 0)
	metadataArray := make([]cadence.Value, 0)

	userAddrC := cadence.NewAddress(flow.HexToAddress(addr.String()))
	require.NoError(t, err)

	recipientArray = append(recipientArray, userAddrC)

	mintTx.AddArgument(cadence.NewArray(recipientArray))

	// Construct arguments with sample metadata
	metadata := lib.NewMetadataV1("test-name", "0.0.1",
		"test-description",
		"https://matrixworld.org/home",
		"https://matrixworld.org/static/media/matrixLogo.9fdc86f0.svg",
		"{}",
		"https://d2yoccx42eml7e.cloudfront.net/metadata/default.mp4",
		"Blocto&Matrixworld").ToCadenceDictionary()
	metadataArray = append(metadataArray, metadata)

	mintTx.AddArgument(cadence.NewArray(metadataArray))

	royaltyAddressC := cadence.NewAddress(flow.HexToAddress(assetsAddr.Hex()))
	royaltyFeeC, err := cadence.NewUFix64("0.05")
	require.NoError(t, err)

	mintTx.AddArgument(royaltyAddressC)
	mintTx.AddArgument(royaltyFeeC)

	mintR, err := e.SignAndExecTrans(mintTx,
		[]flow.Address{e.ServiceKey().Address, assetsAddr},
		[]crypto.Signer{e.ServiceKey().Signer(), signer})
	require.NoError(t, err)
	require.False(t, mintR.Reverted(), mintR.Error)

	t.Logf("Mint: %s", mintR.Events)

}
