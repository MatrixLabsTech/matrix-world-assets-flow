package mint

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts"
	"github.com/MatrixLabsTech/matrix-world-assets-flow/packages/sdk/go/contracts/lib"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
)

const (
	MintFile                           = "mint.cdc"
	defaultNonFungibleTokenAddress     = "\"../contracts/lib/NonFungibleToken.cdc\""
	defaultMatrixWorldAssetsNFTAddress = "\"../contracts/MatrixWorldAssetsNFT.cdc\""
)

func GenerateMintNFTTransaction(nftAddr, mwAssetsNFTAddr string) []byte {
	trans, err := ioutil.ReadFile(filepath.Join(contracts.GetTransRoot(), MintFile))
	if err != nil {
		panic(err)
	}

	// substitute contracts addresses
	scriptWithAddr := strings.ReplaceAll(string(trans), defaultNonFungibleTokenAddress, "0x"+nftAddr)
	scriptWithAddr = strings.ReplaceAll(scriptWithAddr, defaultMatrixWorldAssetsNFTAddress, "0x"+mwAssetsNFTAddr)

	return []byte(scriptWithAddr)
}

func BuildMintNFTsTransactionV1(script []byte,
	recipients []string,
	metadata []lib.MetadataV1,
	signerAdddress flow.Address,
	singerKeyIndex int,
	signer crypto.Signer,
	flowClient *client.Client,
	ctx context.Context,
	royaltyAddress string,
	royaltyFee string,
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

	recipientArray := make([]cadence.Value, 0)
	metadataArray := make([]cadence.Value, 0)

	for _, recipient := range recipients {
		recipientArray = append(recipientArray, cadence.NewAddress(flow.HexToAddress(recipient)))
	}

	for _, metadata := range metadata {
		metadataArray = append(metadataArray, metadata.ToCadenceDictionary())
	}

	tx.AddArgument(cadence.NewArray(recipientArray))
	tx.AddArgument(cadence.NewArray(metadataArray))

	royaltyAddressC := cadence.NewAddress(flow.HexToAddress(royaltyAddress))
	royaltyFeeC, err := cadence.NewUFix64(royaltyFee)

	tx.AddArgument(royaltyAddressC)
	tx.AddArgument(royaltyFeeC)
	if err := tx.SignEnvelope(signerAdddress, singerKeyIndex, signer); err != nil {
		panic(err)
	}

	return tx, nil

}
