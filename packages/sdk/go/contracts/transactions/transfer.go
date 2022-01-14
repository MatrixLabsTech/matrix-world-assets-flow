package transactions

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
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


func BuildTransferNFTTransaction(script []byte,
	recipient string,
  tokenId uint64,
	signerAdddress flow.Address,
  tokenName string,
	singerKeyIndex int,
	signer crypto.Signer,
	flowClient *client.Client,
	ctx context.Context,
) (*flow.Transaction, error) {
	// build transaction
	referenceBlock, err := flowClient.GetLatestBlock(ctx, false)
	if err != nil {
		return nil, err
	}

	proposalAccount, err := flowClient.GetAccountAtLatestBlock(ctx, signerAdddress)
	if err != nil {
		return nil, err
	}

	tx := flow.NewTransaction().
		SetScript([]byte(script)).
		SetGasLimit(9999).
		SetProposalKey(signerAdddress, singerKeyIndex, proposalAccount.Keys[singerKeyIndex].SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(signerAdddress).
		AddAuthorizer(signerAdddress)

  recipientC := cadence.NewAddress(flow.HexToAddress(recipient))
  tokenIdC := cadence.NewUInt64(tokenId)
  typeC, err := cadence.NewString(tokenName)
  if err != nil {
    return nil, err
  }

	tx.AddArgument(tokenIdC)
	tx.AddArgument(recipientC)
	tx.AddArgument(typeC)

	if err := tx.SignEnvelope(signerAdddress, singerKeyIndex, signer); err != nil {
		panic(err)
	}

	return tx, nil

}
