package agreement

import "time"

type SaleAgreement struct {
    ID                 int
    Seller             string
    Buyer              string
    VehicleMake        string
    RegistrationNumber string
    PurchasePrice      int
    IsCompleted        bool
    DateCreated        time.Time
}

func NewSaleAgreement(seller, buyer, make, regNum string, price int) SaleAgreement {
    return SaleAgreement{
        Seller:             seller,
        Buyer:              buyer,
        VehicleMake:        make,
        RegistrationNumber: regNum,
        PurchasePrice:      price,
        IsCompleted:        false,
        DateCreated:        time.Now(),
    }
}