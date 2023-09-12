package student

// handler.getUUID()
type Student struct {
	ID    int64  `json:"ID" gorm:"column:ID";AUTO_INCREMENT`
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email;unique" `
}
