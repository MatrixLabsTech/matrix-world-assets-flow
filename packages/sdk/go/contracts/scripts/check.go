package scripts

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	CheckScriptFile       = "check.cdc"
	defaultNonFungibleTokenAddress = "\"../contracts/lib/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

// CheckScript reads and returns the check.cdc script in bytes
func CheckScript(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	fmt.Println("filename: ", getScriptRoot() + CheckScriptFile)
	script, err := ioutil.ReadFile(getScriptRoot() + CheckScriptFile)
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(script), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	fmt.Println("nftAddr:", nftAddr)
	fmt.Println("mwAssetsNFTAddr:", mwAssetsNFTAddr)
	fmt.Println("script:", scriptWithAddr)
	return []byte(scriptWithAddr)
}
