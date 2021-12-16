import NonFungibleToken from "../../../../contracts/core/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// check MatrixWorldAssetsNFT collection is available on given address
pub fun main(address: Address): Bool {
    return getAccount(address)
        .getCapability<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath)
        .check()
}
