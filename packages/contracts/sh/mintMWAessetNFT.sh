ACCOUNT=0x01cf0e2f2f715450

flow transactions send ./cadence/transactions/setup_accounts.cdc --signer test-admin-account
flow scripts execute ./cadence/scripts/check.cdc $ACCOUNT
flow scripts execute ./cadence/scripts/get_ids.cdc $ACCOUNT

echo "minting nft"
metadata=`jq -c . < sh/metadata.json`
echo $metadata
flow transactions send ./cadence/transactions/mint.cdc --log debug --args-json $metadata '[]' --signer test-admin-account

flow scripts execute ./cadence/scripts/check.cdc $ACCOUNT
flow scripts execute ./cadence/scripts/get_ids.cdc $ACCOUNT