import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// transfer MatrixWorldAssetsNFT token with tokenId to given address
transaction(tokenIds: [UInt64], recipients: [Address], types: [String]) {

    let senderCollection: &MatrixWorldAssetsNFT.Collection

    prepare(acct: AuthAccount) {
        self.senderCollection = acct.borrow<&MatrixWorldAssetsNFT.Collection>(from: MatrixWorldAssetsNFT.collectionStoragePath)
            ?? panic("Missing NFT collection on signer account")
    }

    execute {
        var size = tokenIds.length
        if (size != recipients.length) {
            panic("tokenIds and recipients must be of same length")
        }
        if (size != types.length) {
            panic("tokenIds and types must be of same length")
        }

        while size > 0 {
            let tokenId = tokenIds[size - 1]
            let recipient = recipients[size - 1]
            let type = types[size - 1]
            
            // checkType
            let tokenCollection = self.senderCollection.getMetadata(id: tokenId)["collection"]
            if (tokenCollection != type) {
                panic("tokenId does not belong to the given type")
            }

            // tranfer token
            let token <- self.senderCollection.withdraw(withdrawID: tokenId)
            let receiverProvider = getAccount(recipient).getCapability<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath)
            let receiver = receiverProvider.borrow() ?? panic("Missing NFT collection on receiver account")
            receiver.deposit(token: <- token)
            size = size - 1
        }

    }
}

