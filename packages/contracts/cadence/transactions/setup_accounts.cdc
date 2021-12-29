import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"


// Setup storage for MatrixWorldAssetsNFT on signer account
transaction {
    prepare(acct: AuthAccount) {
        if acct.borrow<&MatrixWorldAssetsNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath) == nil {
            let collection <- MatrixWorldAssetsNFT.createEmptyCollection() as! @MatrixWorldAssetsNFT.Collection
            acct.save(<-collection, to: MatrixWorldAssetsNFT.collectionStoragePath)
            acct.link<&MatrixWorldAssetsNFT.Collection{NonFungibleToken.CollectionPublic, NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath, target: MatrixWorldAssetsNFT.collectionStoragePath)
        }
    }
}
