package student

import "time"

type StudentCourse struct {
    StudentID int `gorm:"primaryKey"` 
    CourseID  int `gorm:"primaryKey"` 
    EnrollmentDate time.Time 
    Student Student
    Course Course
}
