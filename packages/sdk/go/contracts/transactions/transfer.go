package transactions

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
)

func GetTransferAssetCadenceCode(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	script, err := ioutil.ReadFile(filepath.Join(contracts.GetTransRoot(), TransferAssetsTrans))
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(script), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}


