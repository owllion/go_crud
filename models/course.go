package student

type Course struct {
	ID          int          `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Code        string       `json:"code" gorm:"column:code"`
	Name        string       `json:"name" gorm:"column:name"`
	Teacher     string       `json:"teacher" gorm:"column:teacher"`
	Credits     int          `json:"credits" gorm:"column:credits"`
	Enrollments []Enrollment `json:"enrollments" gorm:"foreignKey:CourseRef"`
}
