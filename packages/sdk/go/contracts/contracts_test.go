// Package test performs tests on the contracts method
package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonFungibleToken(t *testing.T) {
	e := NewEmulator()

	code, err := GenerateNonFungibleToken()
	assert.NoError(t, err)
	addr, err := e.Deploy(code, "NonFungibleToken")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	t.Logf("NonFungibleToken address: %s", addr)
}

func TestLicensedNFT(t *testing.T) {
	e := NewEmulator()

	code, err := GenerateLicensedNFT()
	assert.NoError(t, err)
	addr, err := e.Deploy(code, "LicensedNFT")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	t.Logf("LicensedNFT address: %s", addr)
}

func TestMatrixWorldAssetsNFT(t *testing.T) {
	e := NewEmulator()

	nftCode, err := GenerateNonFungibleToken()
	assert.NoError(t, err)
	nftAddr, err := e.Deploy(nftCode, "NonFungibleToken")
	licensedNFTCode, err := GenerateLicensedNFT()
	assert.NoError(t, err)
	licensedNFTAddr, err := e.Deploy(licensedNFTCode, "LicensedNFT")

	code, err := GenerateMatrixWorldAssetsNFT("0x"+nftAddr.String(), "0x"+licensedNFTAddr.String(), getContractRoot())
	assert.NoError(t, err)
	addr, err := e.Deploy(code, "MatrixWorldAssetsNFT")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	t.Logf("MatrixWorldAssetsNFT address: %s", addr)
}