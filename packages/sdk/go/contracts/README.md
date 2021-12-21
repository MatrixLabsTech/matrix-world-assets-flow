# Matrix World Assets NFT Flow SDK

This repository contains the Go SDK for the MatrixWorldAssetsNFT contracts.

## Build and Test

```bash
make
```

### Test Verbosely

```bash
make test
```

## SDK

- `contracts.go`: This contains the code to deploy MatrixWorldAssetsNFT contracts and an emulator to execute scripts and transactions
- `./scripts/`: This package generates all the scripts code from Cadence source file (.cdc)
- `./transactions/`: This package generates all the transactions code from Cadence source file (.cdc)

## Examples

### Contract Deployments

#### generate contract codes

```go
nftCode, err := GenerateNonFungibleToken()

licensedNFTCode, err := GenerateLicensedNFT()

mwNFTCode, err := GenerateMatrixWorldAssetsNFT(
  "0x"+nftAddr.String(),
  "0x"+licensedNFTAddr.String(),
  getContractRoot()
)
```

#### deploy a contract

```go
code, err := GenerateNonFungibleToken()
addr, err := e.Deploy(code, "NonFungibleToken")
```

### Create Accounts

#### create a new user account

```go
accountKeys := test.AccountKeyGenerator()
pk, signer := accountKeys.NewWithSigner()
addr, err := e.CreateAccount([]*flow.AccountKey{pk}, nil)
require.NoError(t, err)
```

### Scripts and Transactions

#### get script codes

```go
ss := scripts.GetCheckScript(nftAddr.String(), mwNFTAddr.String())
```

#### execute script

```go
sr, err := e.ExecuteScript(ss,[][]byte{json.MustEncode(cadence.Address(mwNFTAddr))})
```

#### get transaction codes

```go
ts := transactions.SetupAccounts(nftAddr.String(), mwNFTAddr.String())
```

#### create a transaction

```go
tx := e.CreateTrans(ts, e.ServiceKey().Address)
```

#### execute a transaction

```go
tr, err := e.SignAndExecTrans(
  tx,
  []flow.Address{e.ServiceKey().Address, addr},
  []crypto.Signer{e.ServiceKey().Signer(), signer}
)
```

## Reference

[Flow Go SDK](https://github.com/onflow/flow-go-sdk)

[Flow Go SDK API](https://pkg.go.dev/github.com/onflow/flow-go-sdk)

[NBA Top Shot SDK](https://github.com/dapperlabs/nba-smart-contracts/blob/master/lib/go/test/examples.go)
