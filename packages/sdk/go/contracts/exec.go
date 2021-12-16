package contracts

import (
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// ExecTrans submits a transaction and checks if it fails or not.
func ExecTrans(
	t *testing.T,
	b *emulator.Blockchain,
	tx *flow.Transaction,
	shouldRevert bool,
) {
	// submit the signed transaction
	err := b.AddTransaction(*tx)
	require.NoError(t, err)

	result, err := b.ExecuteNextTransaction()
	require.NoError(t, err)

	if shouldRevert {
		assert.True(t, result.Reverted())
	} else {
		if !assert.True(t, result.Succeeded()) {
			t.Log(result.Error.Error())
		}
	}

	_, err = b.CommitBlock()
	assert.NoError(t, err)
}

// SignAndExecTrans signs a transaction with an array of signers and adds their signatures to the transaction
// before submitting it to the emulator.
//
// If the private keys do not match up with the addresses, the transaction will not succeed.
//
// The shouldRevert parameter indicates whether the transaction should fail or not.
//
// This function asserts the correct result and commits the block if it passed.
func SignAndExecTrans(
	t *testing.T,
	b *emulator.Blockchain,
	tx *flow.Transaction,
	signerAddresses []flow.Address,
	signers []crypto.Signer,
	shouldRevert bool,
) {
	// sign transaction with each signer
	for i := len(signerAddresses) - 1; i >= 0; i-- {
		signerAddress := signerAddresses[i]
		signer := signers[i]

		if i == 0 {
			err := tx.SignEnvelope(signerAddress, 0, signer)
			assert.NoError(t, err)
		} else {
			err := tx.SignPayload(signerAddress, 0, signer)
			assert.NoError(t, err)
		}
	}

	ExecTrans(t, b, tx, shouldRevert)
}