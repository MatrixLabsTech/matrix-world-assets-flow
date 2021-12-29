import NonFungibleToken from "../contracts/lib/NonFungibleToken.cdc"
import MatrixWorldAssetsNFT from "../contracts/MatrixWorldAssetsNFT.cdc"

// Mint MatrixWorldAssetsNFT token to recipient acct
transaction(recipients: [Address], metadata: [{String: String}], royaltyAddress: Address, royaltyFee: UFix64 ) {
    let minter: &MatrixWorldAssetsNFT.Minter

    prepare(acct: AuthAccount) {

        self.minter = acct.borrow<&MatrixWorldAssetsNFT.Minter>(from: MatrixWorldAssetsNFT.minterStoragePath)!;
    }

    execute {
        let ros = [MatrixWorldAssetsNFT.Royalty(address: royaltyAddress, fee: royaltyFee)]
        var size = recipients.length
        while size > 0 {
            let recipient = getAccount(recipients[size - 1])
            let metadata = metadata[size - 1]
            let receiver = recipient.getCapability<&{NonFungibleToken.Receiver}>(MatrixWorldAssetsNFT.collectionPublicPath)
            self.minter.mintTo(creator: receiver, metadata: metadata, royalties: ros)
            size = size - 1
        }
    }
}
