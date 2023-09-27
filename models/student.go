package student

import "time"

type Student struct {
	ID            int          `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	StudentID     string       `json:"studentId" gorm:"column:studentId"`
	Name          string       `json:"name" gorm:"column:name"`
	BirthDate     time.Time    `json:"birthDate" gorm:"column:birthDate"`
	AdmissionYear int          `json:"admissionYear" gorm:"column:admissionYear"`
	Enrollments   []Enrollment `json:"enrollments" gorm:"foreignKey:StudentRef"`
}
