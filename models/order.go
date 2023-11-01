package student

import "time"


type Order struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	TotalPrice float64  `gorm:"type:decimal(10,2);column:total_price" json:"total_price"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
}



func (Order) TableName() string {
	return "enrollment.order"
}