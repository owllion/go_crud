package scRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"strconv"

	"github.com/gin-gonic/gin"
)



func DeleteStudentCourse(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	studentCourse := student.StudentCourse{}
	g.Ctx.ShouldBind(&studentCourse)

	studentId := ctx.Query("studentId")
	courseId := ctx.Query("courseId")

	if studentId != "" && courseId != "" {
		sID, err1 := strconv.Atoi(studentId)
		cID, err2 := strconv.Atoi(courseId)

		if err1 != nil || err2 != nil {
			g.SendResponse(400, "Invalid student or course id", nil)
			return
		}

		result := db.MysqlDB.Debug().Where("`StudentID` = ? AND `CourseID` = ?", sID, cID).Delete(&studentCourse)

		if result.RowsAffected == 0 { 
			g.SendResponse(404, "The association between student and course does not exist", nil)
			return
		}

		if result.Error != nil {
			g.SendResponse(500,"Fail to delete the association",result.Error.Error())
			return
		}
		g.SendResponse(200, "Deleted the association successfully", nil)
		return
	}
	g.SendResponse(400, "Please provide the student ID and course ID.", nil)
}
