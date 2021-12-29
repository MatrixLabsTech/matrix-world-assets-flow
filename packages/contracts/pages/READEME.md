# Local TEST

### setup account - init collection

```bash
flow transactions send cadence/transactions/setup_accounts.cdc
```

### get nft ids

```bash
flow scripts execute cadence/scripts/get_ids.cdc f8d6e0586b0a20c7
```

### get nft metadata 

```bash
flow scripts execute cadence/scripts/get_metadata.cdc f8d6e0586b0a20c7 0
```

### mint test ids using admin account

```bash
flow transactions send cadence/transactions/mint.cdc '[0xf8d6e0586b0a20c7]' '[{"version":"1"}]' 01cf0e2f2f715450 0.05 --signer test-admin-account
```

### check collection

```bash
flow scripts execute cadence/scripts/check.cdc 01cf0e2f2f715450  
```
