import FungibleToken from "../../contracts/lib/FungibleToken.cdc"
import FUSD from "../../contracts/lib/FUSD.cdc"

transaction {
  prepare(signer: AuthAccount) {

    if(signer.borrow<&FUSD.Vault>(from: /storage/fusdVault) != nil) {
      return
    }

    signer.save(<-FUSD.createEmptyVault(), to: /storage/fusdVault)

    signer.link<&FUSD.Vault{FungibleToken.Receiver}>(
      /public/fusdReceiver,
      target: /storage/fusdVault
    )

    signer.link<&FUSD.Vault{FungibleToken.Balance}>(
      /public/fusdBalance,
      target: /storage/fusdVault
    )
  }
}
