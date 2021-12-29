import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// Take MatrixWorldAssetsNFT ids by account address
pub fun main(address: Address): [UInt64]? {
    let collection = getAccount(address)
        .getCapability(MatrixWorldAssetsNFT.collectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver}>()
        ?? panic("NFT Collection not found")
    return collection.getIDs()
}
