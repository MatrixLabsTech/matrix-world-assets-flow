// Package contracts generates a new MatrixWorldAssetsNFT contract
package contracts

import (
	"path/filepath"

	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-emulator/types"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"

	"io/ioutil"
	"strings"
)

const (
	matrixWorldAssetsNFTFile     = "MatrixWorldAssetsNFT.cdc"
	nonFungibleTokenContractsURL = "https://raw.githubusercontent.com/onflow/flow-nft/master/contracts/NonFungibleToken.cdc"
	licensedNFTInterfaceURL      = "https://raw.githubusercontent.com/rarible/flow-contracts/main/contracts/LicensedNFT.cdc"

	defaultNonFungibleTokenAddress = "\"lib/NonFungibleToken.cdc\""
	defaultLicensedNFT             = "\"LicensedNFT.cdc\""
)

// GenerateMatrixWorldAssetsNFT returns a copy of the MatrixWorldAssetsNFT contract.
// The contract address is replaced with the given nftAddr and licensedNftAddr.
func GenerateMatrixWorldAssetsNFT(nftAddr, licensedNftAddr, contractRoot string) (code contractCode, err error) {
	// read the contract file as string
	contractFile, err := ioutil.ReadFile(filepath.Join(contractRoot, matrixWorldAssetsNFTFile))
	if err != nil {
		return nil, err
	}

	// substitute contracts addresses
	codeWithAddr := strings.ReplaceAll(string(contractFile), defaultNonFungibleTokenAddress, nftAddr)
	codeWithAddr = strings.ReplaceAll(codeWithAddr, defaultLicensedNFT, licensedNftAddr)

	return contractCode(codeWithAddr), nil
}

// GenerateNonFungibleToken returns a copy of the Flow NunFungibleToken contract.
func GenerateNonFungibleToken() (code contractCode, err error) {
	nftCode, err := downloadFile(nonFungibleTokenContractsURL)
	if err != nil {
		return nil, err
	}
	return nftCode, nil
}

// GenerateLicensedNFT returns a copy of the LicensedNFT contract.
func GenerateLicensedNFT() (code contractCode, err error) {
	licensedNFTCode, _ := downloadFile(licensedNFTInterfaceURL)
	if err != nil {
		return nil, err
	}
	return licensedNFTCode, nil
}

type contractCode []byte

type flowEmulator struct {
	emulator.Blockchain
}

// NewEmulator returns a flow blockchain emulator for testing.
func NewEmulator(opts ...emulator.Option) *flowEmulator {
	b, err := emulator.NewBlockchain(
		append(
			[]emulator.Option{
				emulator.WithStorageLimitEnabled(false),
			},
			opts...,
		)...,
	)
	if err != nil {
		panic(err)
	}

	return &flowEmulator{*b}
}

// Deploy deploys a contract onto the emulator and returns its address
func (e *flowEmulator) Deploy(code contractCode, contract string) (flow.Address, error) {
	addr, err := e.CreateAccount(nil, []templates.Contract{
		{
			Name:   contract,
			Source: string(code),
		},
	})
	if err != nil {
		return flow.Address{}, err
	}
	_, err = e.CommitBlock()
	if err != nil {
		return flow.Address{}, err
	}
	return addr, err
}

// SignAndExecTrans signs a transaction with an array of signers and adds their signatures to the transaction
// before submitting it to the emulator.
//
// If the private keys do not match up with the addresses, the transaction will not succeed.
//
// The shouldRevert parameter indicates whether the transaction should fail or not.
//
// This function asserts the correct result and commits the block if it passed.
func (e *flowEmulator) SignAndExecTrans(
	tx *flow.Transaction,
	signerAddresses []flow.Address,
	signers []crypto.Signer,
) (*types.TransactionResult, error) {
	// sign transaction with each signer
	for i := len(signerAddresses) - 1; i >= 0; i-- {
		signerAddress := signerAddresses[i]
		signer := signers[i]

		var err error
		if i == 0 {
			err = tx.SignEnvelope(signerAddress, 0, signer)
		} else {
			err = tx.SignPayload(signerAddress, 0, signer)
		}
		if err != nil {
			return nil, err
		}
	}
	return e.ExecTrans(tx)
}

// ExecTrans executes the given transaction on the emulator
func (e *flowEmulator) ExecTrans(tx *flow.Transaction) (*types.TransactionResult, error) {
	// submit the signed transaction
	err := e.AddTransaction(*tx)

	r, err := e.ExecuteNextTransaction()
	if r.Reverted() {
		return r, err
	}

	_, err = e.CommitBlock()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// CreateTrans creates a transaction with the given transaction byte code and authorizer
func (e *flowEmulator) CreateTrans(script []byte, authorizerAddress flow.Address) *flow.Transaction {
	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(flow.DefaultTransactionGasLimit).
		SetProposalKey(e.ServiceKey().Address, e.ServiceKey().Index, e.ServiceKey().SequenceNumber).
		SetPayer(e.ServiceKey().Address).
		AddAuthorizer(authorizerAddress)
	return tx
}
