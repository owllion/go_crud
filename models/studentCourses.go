package student

import "time"

type StudentCourse struct {
    StudentID int `gorm:"primaryKey"` 
    CourseID  int `gorm:"primaryKey"` 
    EnrollmentDate time.Time 
    Student Student `json:"-" gorm:"foreignKey:StudentID;references:ID"`
    Course Course `json: "-" gorm:"foreignKey:CourseID;references:ID"`
}

//NOTE: 只寫 `json:"-"，完全不會影響最終回傳結果，只會受preload影響
//NOTE:　寫完整的　`json:"-" gorm:"foreignKey:StudentID;references:ID"`，才會真的讓最終結果不回傳此欄位
//NOTE:　gorm:"foreignKey:StudentID;references:ID　->  表示 StudentID 是用來與 Student 表的 ID 欄位建立關聯的外部鍵，

//TODO: 1-1 一個學生選一門課，學生id是pk，課程id是fk(和其他表建立關聯)，確保課程是真的存在

//TODO: m-m 一個學生選多門，一門可被多人選，中間表pk來自學生&課程pk(id)，同時也需寫fk去做關聯 -> user_id = Column(String(80), ForeignKey("user.id"), primary_key=True, nullable=False) ///// coupon_id = Column(String(80), ForeignKey("coupon.id"), primary_key=True, nullable=False)，可以看到sqlAlchemy寫法和gorm非常類似，感覺是只有差在py的要自己寫 "coupon.id"(關聯的table & 要關聯的欄位，相當於 foreignKey:StudentID;references:ID" )

