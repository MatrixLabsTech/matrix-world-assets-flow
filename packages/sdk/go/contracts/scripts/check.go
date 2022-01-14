package scripts

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
)

const (
	CheckScriptFile                    = "check.cdc"
	defaultNonFungibleTokenAddress     = "\"../contracts/lib/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

// GetCheckScript reads and returns the check.cdc script in bytes
func GetCheckScript(nftAddr, mwAssetsNFTAddr string) []byte {
	// read the script file as string
	script, err := ioutil.ReadFile(filepath.Join(contracts.GetScriptRoot(), CheckScriptFile))
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(script), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}

func CheckUserCollection(ctx context.Context, script []byte, userAddress string, flowClient *client.Client) (bool, error) {
  args := []cadence.Value{cadence.NewAddress(flow.HexToAddress(userAddress))}
  value, err := flowClient.ExecuteScriptAtLatestBlock(ctx, script, args)
  if err != nil {
    return false, err
  }
  return value.ToGoValue().(bool), nil
}
