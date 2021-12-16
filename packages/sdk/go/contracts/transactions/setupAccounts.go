package transactions

import (
	"io/ioutil"
	"strings"
)

const (
	SetupAccountsFile       = "setup_accounts.cdc"
	defaultNonFungibleTokenAddress = "\"../contracts/lib/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

// SetupAccounts reads and returns the check.cdc script in bytes
func SetupAccounts(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	trans, err := ioutil.ReadFile(getTransactionsRoot() + SetupAccountsFile)
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(trans), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}
