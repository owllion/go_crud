package courseRoute

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
		g.SendResponse(500, "Fail to add student course", nil)
		return
	}

	g.SendResponse(200,"create student course successfully",nil)

	
}

