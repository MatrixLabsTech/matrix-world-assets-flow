// Package contracts generates a new MatrixWorldAssetsNFT contract
package contracts

import (
	"github.com/onflow/flow-emulator"
	"github.com/onflow/flow-emulator/types"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"

	"io/ioutil"
	"strings"
)

const (
	MatrixWorldAssetsNFTFile       = "MatrixWorldAssetsNFT.cdc"
	defaultNonFungibleTokenAddress = "\"lib/NonFungibleToken.cdc\""
	defaultLicensedNFT             = "\"LicensedNFT.cdc\""
)

// GenerateMatrixWorldAssetsNFT returns a copy of the MatrixWorldAssetsNFT contract.
// The contract address is replaced with the given nftAddr and licensedNftAddr.
func GenerateMatrixWorldAssetsNFT(nftAddr, licensedNftAddr, contractRoot string) (code ContractCode, err error) {
	// read the contract file as string
	contractFile, err := ioutil.ReadFile(contractRoot + MatrixWorldAssetsNFTFile)
	if err != nil {
		return nil, err
	}

	// substitute contracts addresses
	codeWithAddr := strings.ReplaceAll(string(contractFile), defaultNonFungibleTokenAddress, nftAddr)
	codeWithAddr = strings.ReplaceAll(codeWithAddr, defaultLicensedNFT, licensedNftAddr)

	return ContractCode(codeWithAddr), nil
}

func GenerateNonFungibleToken() (code ContractCode, err error) {
	nftCode, err := DownloadFile(NonFungibleTokenContractsBaseURL + NonFungibleTokenInterfaceFile)
	if err != nil {
		return nil, err
	}
	return ContractCode(nftCode), nil
}

func GenerateLicensedNFT() (code ContractCode, err error) {
	licensedNFTCode, _ := DownloadFile(LicensedNFTInterfaceURL)
	if err != nil {
		return nil, err
	}
	return ContractCode(licensedNFTCode), nil
}

type ContractCode []byte

type flowEmulator struct {
	emulator.Blockchain
}

// NewEmulator returns an emulator blockchain for testing.
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

func (e *flowEmulator) Deploy(code ContractCode, contract string) (flow.Address, error) {
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
	shouldRevert bool,
) (*types.TransactionResult, error){
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


