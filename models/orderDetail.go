package model

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ID           uint    `gorm:"column:id" json:"id"`
	OrderID      uint    `gorm:"column:orderid" json:"orderid"`
	ProductID    uint    `gorm:"column:productid" json:"productid"`
	Quantity     uint    `gorm:"column:quantity" json:"quantity"`
	Price        float64 `gorm:"column:price" json:"price"`
	Currency     string  `gorm:"column:currencytype;references:CurrencyType" json:"currency"`
	ExchangeRate string  `gorm:"column:exchangerate;references:exchangerate" json:"exchangerate"`
}
