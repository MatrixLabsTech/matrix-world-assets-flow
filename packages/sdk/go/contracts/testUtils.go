package contracts

import (
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"

	emulator "github.com/onflow/flow-emulator"
)

//func DeployMatrixWorldAssetsNFT(t *testing.T, contractRoot string) (b *emulator.Blockchain, assetsNFTAddr flow.Address) {
//	b = NewEmulator()
//
//	// Should be able to deploy the NFT contract
//	// as a new account with no keys.
//	nftCode, _ := DownloadFile(NonFungibleTokenContractsBaseURL + NonFungibleTokenInterfaceFile)
//	nftAddr, err := b.CreateAccount(nil, []sdkTemplates.Contract{
//		{
//			Name:   "NonFungibleToken",
//			Source: string(nftCode),
//		},
//	})
//	t.Logf("NFT address: %s", nftAddr)
//	if !assert.NoError(t, err) {
//		t.Log(err.Error())
//	}
//	_, err = b.CommitBlock()
//	assert.NoError(t, err)
//
//	// deploy LicensedNFTInterface
//	licensedNFTCode, _ := DownloadFile(LicensedNFTInterfaceURL)
//	licensedNFTAddr, err := b.CreateAccount(nil, []sdkTemplates.Contract{
//		{
//			Name:   "LicensedNFT",
//			Source: string(licensedNFTCode),
//		},
//	})
//	t.Logf("LicensedNFT address: %s", licensedNFTAddr)
//	if !assert.NoError(t, err) {
//		t.Log(err.Error())
//	}
//	_, err = b.CommitBlock()
//	assert.NoError(t, err)
//
//	// Should be able to deploy AssetNFT contract
//	// as a new account with no keys.
//	assetsNFTCode, _ := GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(),
//		getContractRoot())
//	t.Logf("assetsNFTCode contract: %s", assetsNFTCode)
//	assetsNFTAddr, err = b.CreateAccount(nil, []sdkTemplates.Contract{
//		{
//			Name:   "MatrixWorldAssetsNFT",
//			Source: string(assetsNFTCode),
//		},
//	})
//	t.Logf("MatrixWorldAssetsNFT address: %s", assetsNFTAddr)
//	if !assert.NoError(t, err) {
//		t.Log(err.Error())
//	}
//	_, err = b.CommitBlock()
//	assert.NoError(t, err)
//	return
//}

// ReadFile reads a file from the file system
func ReadFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return contents
}

// DownloadFile will download an url a byte slice
func DownloadFile(url string) ([]byte, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func createTxWithTemplateAndAuthorizer(
	b *emulator.Blockchain,
	script []byte,
	authorizerAddress flow.Address,
) *flow.Transaction {

	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(9999).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(authorizerAddress)

	return tx
}



func readFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return contents
}

// ExecuteScriptAndCheckShouldFail executes a script and checks to make sure
// that it succeeded
func ExecuteScriptAndCheckShouldFail(t *testing.T, b *emulator.Blockchain, script []byte, shouldRevert bool) {
	result, err := b.ExecuteScript(script, nil)
	if err != nil {
		t.Log(string(script))
	}
	require.NoError(t, err)
	if shouldRevert {
		assert.True(t, result.Reverted())
	} else {
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	}
}

// CadenceUFix64 returns a UFix64 value
func CadenceUFix64(value string) cadence.Value {
	newValue, err := cadence.NewUFix64(value)

	if err != nil {
		panic(err)
	}

	return newValue
}

// CadenceString returns a string value from a string representation
func CadenceString(value string) cadence.Value {
	newValue, err := cadence.NewString(value)

	if err != nil {
		panic(err)
	}

	return newValue
}

func bytesToCadenceArray(b []byte) cadence.Array {
	values := make([]cadence.Value, len(b))

	for i, v := range b {
		values[i] = cadence.NewUInt8(v)
	}

	return cadence.NewArray(values)
}

// assertEqual asserts that two objects are equal.
//
//    assertEqual(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
//
func assertEqual(t *testing.T, expected, actual interface{}) bool {

	if assert.ObjectsAreEqual(expected, actual) {
		return true
	}

	message := fmt.Sprintf(
		"Not equal: \nexpected: %s\nactual  : %s",
		expected,
		actual,
	)

	return assert.Fail(t, message)
}


