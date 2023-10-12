package student

import "time"

type Student struct {
	ID            int          `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	StudentID     string       `json:"studentId" gorm:"column:student_id"`
	Name          string       `json:"name" gorm:"column:name"`
	BirthDate     time.Time    `json:"birthDate" gorm:"column:birth_date"`
	AdmissionYear int          `json:"admissionYear" gorm:"column:admission_year"`
	Courses []Course `json:"courses" gorm:"many2many:student_courses;"`
	//不會存到db
}


func (Student) TableName() string {
	return "enrollment.student"
}