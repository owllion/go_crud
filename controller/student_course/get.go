package scRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)


func GetStudentCourses(ctx *gin.Context) {
	g := handler.GinContext{Ctx:ctx}
	students := []student.StudentCourse{}

	if studentId := ctx.Query("studentId"); studentId != "" {

		result := db.MysqlDB.Debug().Find(&students)

		if result.Error != nil {
			g.SendResponse(500, "Fail to get student's courses", nil)
			return
		}

		g.SendResponse(200,"get student's courses successfully",students)
		return 
	}

	g.SendResponse(400, "Please provide the student ID and course ID.", nil)


	
}
