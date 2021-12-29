import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

pub fun main(address: Address, id: UInt64): &AnyResource {
    let collection = getAccount(address)
        .getCapability(MatrixWorldAssetsNFT.collectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>() ?? panic("NFT Collection not found")
    let nft = collection.borrowNFT(id: id)
    return nft
}
