package contracts

import (
	"io/ioutil"
	"strings"
)

const (
	MatrixWorldAssetsNFTFile       = "MatrixWorldAssetsNFT.cdc"
	defaultNonFungibleTokenAddress = "\"lib/NonFungibleToken.cdc\""
	defaultLicensedNFT             = "\"LicensedNFT.cdc\""
)

func GenerateTopShotContract(nftAddr string, licensedNftAddr string, contractRoot string) []byte {
	// root to './' if empty
	if contractRoot == "" {
		contractRoot = "./"
	}

	// read the contract file as string
	contractFile, err := ioutil.ReadFile(contractRoot + MatrixWorldAssetsNFTFile)
	if err != nil {
		panic(err)
	}

	// convert to string
	codeWithAddr := strings.ReplaceAll(string(contractFile), defaultNonFungibleTokenAddress, nftAddr)
	codeWithAddr = strings.ReplaceAll(codeWithAddr, defaultLicensedNFT, licensedNftAddr)

	return []byte(codeWithAddr)
}
