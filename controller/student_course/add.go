package scRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)

func CreateStudentCourse(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	req := []student.StudentCourse{}
	g.Ctx.ShouldBind(&req)
	
	result := db.MysqlDB.Debug().Create(&req)


	if result.Error != nil {
		g.SendResponse(500, "Fail to create student course", nil)
		return
	}

	g.SendResponse(200,"Create student successfully",nil)

	
}

