import NonFungibleToken from "../../../contracts/core/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// Burn MatrixWorldAssetsNFT on signer account by tokenId
transaction(tokenId: UInt64) {
    prepare(account: AuthAccount) {
        let collection = account.borrow<&RaribleNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath)!
        destroy collection.withdraw(withdrawID: tokenId)
    }
}