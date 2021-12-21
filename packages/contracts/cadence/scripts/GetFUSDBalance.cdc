import FungibleToken from "../contracts/lib/FungibleToken.cdc"
import FUSD from "../contracts/lib/FUSD.cdc"

pub fun main(address: Address): UFix64? {
  let account = getAccount(address)

  if let vaultRef = account.getCapability(/public/fusdBalance).borrow<&FUSD.Vault{FungibleToken.Balance}>() {
    return vaultRef.balance
  }
  return nil

}
