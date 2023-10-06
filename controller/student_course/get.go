package scRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetStudentCourses(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	studentCourses := []student.StudentCourse{}

	studentId := ctx.Query("studentId")

	if studentId != ""  {
		sID, err1 := strconv.Atoi(studentId)

		if err1 != nil {
			g.SendResponse(400, "Invalid student id", nil)
			return
		}

		/*NOTE 
		不管這邊有無Preload Student，最後都會return一個Student obj，
		
		有 Preload -> 加載實際資料 / 無preload -> 給全都是空值的struct
		
		如果完全不想要他在json裡面出現，就去struct的該欄位加上 json:"-" 即可
		*/
		 result := db.MysqlDB.Debug().Preload("Student").Preload("Course").Where("`student_id` = ?", sID).Find(&studentCourses)

		if result.Error != nil {
			g.SendResponse(500,"Fail to find the associated data",result.Error.Error())
			return
		}

		if result.RowsAffected == 0 { 
			g.SendResponse(404, "The association between student and course does not exist", nil)
			return
		}
		
		g.SendResponse(200, "Find the associated courses successfully", studentCourses)
		return
	}

	g.SendResponse(400, "Please provide the student ID", nil)
}
