// Simple fee manager
//
pub contract AssetFee {

    pub let commonFeeManagerStoragePath: StoragePath

    pub event SellerFeeChanged(value: UFix64)
    pub event BuyerFeeChanged(value: UFix64)
    pub event FeeAddressUpdated(label: String, address: Address)

    access(self) var feeAddresses: {String:Address}

    // Seller fee [0..1)
    pub var sellerFee: UFix64

    // BuyerFee fee [0..1)
    pub var buyerFee: UFix64

    pub resource Manager {
        pub fun setSellerFee(_ fee: UFix64) {
            AssetFee.sellerFee = fee
            emit SellerFeeChanged(value: AssetFee.sellerFee)
        }

        pub fun setBuyerFee(_ fee: UFix64) {
            AssetFee.buyerFee = fee
            emit BuyerFeeChanged(value: AssetFee.buyerFee)
        }

        pub fun setFeeAddress(_ label: String, address: Address) {
            AssetFee.feeAddresses[label] = address
            emit FeeAddressUpdated(label: label, address: address)
        }
    }

    init() {
        self.sellerFee = 0.025
        emit SellerFeeChanged(value: AssetFee.sellerFee)
        self.buyerFee = 0.025
        emit BuyerFeeChanged(value: AssetFee.buyerFee)

        self.feeAddresses = {}

        self.commonFeeManagerStoragePath = /storage/commonFeeManager
        self.account.save(<- create Manager(), to: self.commonFeeManagerStoragePath)
    }

    pub fun feeAddress(): Address {
        return self.feeAddresses["asset"] ?? self.account.address
    }

    pub fun feeAddressByName(_ label: String): Address {
        return self.feeAddresses[label] ?? self.account.address
    }

    pub fun addressMap(): {String:Address} {
        return AssetFee.feeAddresses
    }
}
