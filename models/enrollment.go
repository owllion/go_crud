package student

type Enrollment struct {
	ID         int     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	StudentRef int     `json:"studentRef" gorm:"column:studentRef"`
	CourseRef  int     `json:"courseRef" gorm:"column:courseRef"`
	Semester   string  `json:"semester" gorm:"column:semester"`
	Grade      float64 `json:"grade" gorm:"column:grade"`
}