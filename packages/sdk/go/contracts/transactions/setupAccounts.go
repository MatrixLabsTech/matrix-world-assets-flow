package transactions

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
)

const (
	SetupAccountsFile                  = "setup_accounts.cdc"
	defaultNonFungibleTokenAddress     = "\"../contracts/lib/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

// SetupAccounts reads and returns the setup_accounts.cdc transaction file in bytes
func SetupAccounts(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	trans, err := ioutil.ReadFile(filepath.Join(contracts.GetTransRoot(), SetupAccountsFile))
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(trans), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}
