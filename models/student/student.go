package student

type Student struct {
	ID    int64  `json:"ID" gorm:"column:ID"`
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}