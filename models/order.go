package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint          `gorm:"column:id" json:"id"`
	OrderNumber  string        `gorm:"column:ordernumber" json:"ordernumber"`
	CustomerID   uint          `gorm:"column:customerid" json:"customerid"`
	TotalPrice   float64       `gorm:"column:totalprice" json:"totalprice"`
	CurrencyType string        `gorm:"column:currencytype" json:"currencytype"`
	ExchangeRate float64       `gorm:"column:exchangerate" json:"exchangerate"`
	Status       string        `gorm:"column:status" json:"status"`
	CreatedAt    time.Time     `gorm:"column:createdat" json:"createdat"`
	UpdatedAt    time.Time     `gorm:"column:updatedat" json:"updatedat"`
	OrderDetail  []OrderDetail `json:"orderDetail"`
}

// type Order struct {
// 	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
// 	TotalPrice float64   `gorm:"type:decimal(10,2);column:total_price" json:"total_price"`
// 	CreatedAt  time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
// }

func (Order) TableName() string {
	return "enrollment.order"
}
