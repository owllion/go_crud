package student

import "time"

type Student struct {
	ID            int          `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	StudentID     string       `json:"studentId" gorm:"column:student_id"`
	Name          string       `json:"name" gorm:"column:name"`
	BirthDate     time.Time    `json:"birthDate" gorm:"column:birth_date"`
	AdmissionYear int          `json:"admissionYear" gorm:"column:admission_year"`
	Courses []Course `json:"courses" gorm:"many2many:enrollment.student_course;;"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`

	//不會存到db
}


func (Student) TableName() string {
	return "enrollment.student"
}