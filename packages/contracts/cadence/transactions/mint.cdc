import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// Mint MatrixWorldAssetsNFT token to signer acct
//
transaction(metadata: MatrixWorldAssetsNFT.Metadata, royalties: [MatrixWorldAssetsNFT.Royalty]) {
    let minter: &MatrixWorldAssetsNFT.Minter
    let receiver: Capability<&{NonFungibleToken.Receiver}>

    prepare(acct: AuthAccount) {
        if acct.borrow<&MatrixWorldAssetsNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath) == nil {
            let collection <- MatrixWorldAssetsNFT.createEmptyCollection() as! @MatrixWorldAssetsNFT.Collection
            acct.save(<- collection, to: MatrixWorldAssetsNFT.collectionStoragePath)
            acct.link<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath, target: MatrixWorldAssetsNFT.collectionStoragePath)
        }

        self.minter = MatrixWorldAssetsNFT.minter()
        self.receiver = acct.getCapability<&{NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath)
    }

    execute {
        log(metadata)
        self.minter.mintTo(creator: self.receiver, metadata: metadata, royalties: royalties)  // FIXME
    }
}
