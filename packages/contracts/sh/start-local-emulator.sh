#!/bin/bash
export TEST_ADMIN_KEY=fbe537d4b7743bce3271633073cdd92ebc9a17f6dbd8e1ff73715c9d2fa9c01651d75d394d9c3c8dfe489eae69879737243a19574a59a6163b9f56f592b4355a

flow accounts create --key $TEST_ADMIN_KEY

flow transactions send ./cadence/transactions/emulator/TransferFlowToken.cdc 100.0 0x01cf0e2f2f715450;

flow project deploy --update
