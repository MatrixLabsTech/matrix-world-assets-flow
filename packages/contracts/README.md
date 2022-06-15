## Deploy Project to Emulator

Start local emulator

`flow emulator`

Deploy all project

`bash ./sh/deploy-to-local-emulator.sh`

## Deploy to testnet

```bash
flow deploy project -n testnet -f flow.json ./envs/flow.testnet.json
```

## Upgrade to mainnet

```bash
flow accounts update-contract MatrixWorldAssetsNFT ./cadence/contracts/MatrixWorldAssetsNFTUpgradeMain.cdc --signer main -f ./flow.json -f ./flow.main.json -n mainnet
```
