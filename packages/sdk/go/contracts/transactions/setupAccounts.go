package transactions

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	SetupAccountsFile                  = "setup_accounts.cdc"
	defaultNonFungibleTokenAddress     = "\"../../../contracts/core/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

// SetupAccounts reads and returns the setup_accounts.cdc transaction file in bytes
func SetupAccounts(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	trans, err := ioutil.ReadFile(filepath.Join(GetRoot(), SetupAccountsFile))
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(trans), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}
