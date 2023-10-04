package studentRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)

func CreateCourse(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	req := []student.Course{}
	g.Ctx.ShouldBind(&req)
	
	result := db.MysqlDB.Debug().Create(&req)


	if result.Error != nil {
		g.SendResponse(500, "新增課程失敗", nil)
		return
	}

	g.SendResponse(200,"新增課程成功",nil)

	
}

