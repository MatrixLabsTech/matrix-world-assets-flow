import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// transfer MatrixWorldAssetsNFT token with tokenId to given address
transaction(tokenId: UInt64, recipient: Address, type: String) {

    let senderCollection: &MatrixWorldAssetsNFT.Collection

    prepare(acct: AuthAccount) {
        self.senderCollection = acct.borrow<&MatrixWorldAssetsNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath)
            ?? panic("Missing NFT collection on signer account")
    }

    execute {
        
        // checkType
        let tokenCollection = self.senderCollection.getMetadata(id: tokenId)["name"]
        if (tokenCollection != type) {
            panic("tokenId does not belong to the given type")
        }

        // tranfer token
        let token <- self.senderCollection.withdraw(withdrawID: tokenId)
        let receiverProvider = getAccount(recipient).getCapability<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath)
        let receiver = receiverProvider.borrow() ?? panic("Missing NFT collection on receiver account")
        receiver.deposit(token: <- token)

    }
}
