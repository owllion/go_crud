package student

import "time"

type AqlRequestJSON struct {
	InspectionLevel string  `json:"inspection_level"`
	CriticalAql     float64 `json:"critical_aql"`
	MajorAql        float64 `json:"major_aql"`
	MinorAql        float64 `json:"minor_aql"`
	QtyFrom         int     `json:"qty_from"`
	QtyTo           int     `json:"qty_to"`
}

type Usersamplingplan struct {
	Id               int64     `json:"id" gorm:"column:id"`
	User_id          int64     `json:"user_id" gorm:"column:user_id"`
	Product_qty      int64     `json:"product_qty" gorm:"column:product_qty"`
	Inspection_level string    `json:"inspection_level" gorm:"column:inspection_level"`
	Sampling_letter  string    `json:"sampling_letter" gorm:"column:sampling_letter"`
	Critical_aql     float64   `json:"critical_aql" gorm:"column:critical_aql"`
	Major_aql        float64   `json:"major_aql" gorm:"column:major_aql"`
	Minor_aql        float64   `json:"minor_aql" gorm:"column:minor_aql"`
	Qty_range_from   int64     `json:"qty_range_from" gorm:"column:qty_range_from"`
	Qty_range_to     int64     `json:"qty_range_to" gorm:"column:qty_range_to"`
	Critical_ac      int64     `json:"critical_ac" gorm:"column:critical_ac"`
	Major_ac         int64     `json:"major_ac" gorm:"column:major_ac"`
	Minor_ac         int64     `json:"minor_ac" gorm:"column:minor_ac"`
	Created_at       time.Time `json:"created_at" gorm:"column:created_at"`
}

type Singleplan struct {
	Id              int64   `json:"id" gorm:"column:id"`
	Ac_num          int64   `json:"ac_num" gorm:"column:ac_num"`
	Re_num          int64   `json:"re_num" gorm:"column:re_num"`
	Sampling_letter string  `json:"sampling_letter" gorm:"column:sampling_letter"`
	Aql             float64 `json:"aql" gorm:"column:aql"`
}