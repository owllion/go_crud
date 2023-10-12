package student

import "time"

type Course struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Code      string    `json:"code" gorm:"column:code"`
	Name      string    `json:"name" gorm:"column:name"`
	Teacher   string    `json:"teacher" gorm:"column:teacher"`
	Credits   int       `json:"credits" gorm:"column:credits"`
	Students  []Student `json:"students" gorm:"many2many:enrollment.student_course;;"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
}

func (Course) TableName() string {
	return "enrollment.course"
}