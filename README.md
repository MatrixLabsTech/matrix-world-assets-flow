# Matrix World Asset NFT

This repository contains the smart contracts for the Matrix World Asset. The Flow NFT was forked from [Rarible flow contracts](https://github.com/rarible/flow-contracts).

The smart contracts are written in [Cadence](https://docs.onflow.org/cadence).

## Addresses

To view the contracts, visit [flow-view-source](https://flow-view-source.com/testnet/account/0xe2f1b000e0203c1d/contract/MatrixWorldAssetsNFT) or using [flowscan-testnet](https://testnet.flowscan.org/account/0xe2f1b000e0203c1d/overview)

| Contract      | Mainnet | Testnet              |
| ------------- | ------- | -------------------- |
| MatrixWorldAssetsNFT      |         | `0xe2f1b000e0203c1d` |


## Deploy

### local emulator

```bash
cd packages/contracts && yarn emulator  # start a flow emulator
yarn run deploy:local # setup accounts and deploy contracts to emulator
```

```bash
yarn run redeploy:local # redeploy all contracts
```

### testnet

```bash
yarn run deploy:testnet
```

```bash
yarn run redeploy:testnet # redeploy all contracts
```

## Structure

```bash
.
├── LICENSE
├── package.json
├── packages
│   ├── contracts   # Cadence projects
│   └── sdk         # TS client
├── README.md
```

## Smart contracts

#### Interface Contracts

[NonFungibleToken](https://docs.onflow.org/core-contracts/non-fungible-token/) : This follows [Flow NFT standard](https://github.com/onflow/flow-nft) which is equivalent to ERC-721 or ERC-1155 on Ethereum. This contract can be directly imported from `0x631e88ae7f1d7c20`when deployed onto testnet or `0x1d7e57aa55817448` when deployed onto mainnet.

`LicensedNFT`: This is a contract interface to be further extended to create a solid NFT contract. It adds royalties to NFT. You can implement this `LicensedNFT` in your
contract (along with [`NonFungibleToken`](https://github.com/onflow/flow-nft)) and your royalties will be taken when trading on [Rarible](https://rarible.com/).

#### Contracts Implementation

`MatrixWorldAssetsNFT`: The Matrix World Asset NFT contract that implements the flow `NonFungibleToken` and `LicensedNFT` contract. 

