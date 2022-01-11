import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// Setup storage for MatrixWorldAssetsNFT on signer account
transaction {
    prepare(acct: AuthAccount) {
        acct.unlink(MatrixWorldAssetsNFT.collectionPublicPath)
        let collection  <- acct.load<@MatrixWorldAssetsNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath)
        destroy collection
    }
}
