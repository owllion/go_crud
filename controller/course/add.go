package courseRoute

import (
	db "practice/database"
	student "practice/models"
	"practice/util"

	"github.com/gin-gonic/gin"
)

func CreateCourse(ctx *gin.Context) {
	g := util.GinContext{Ctx: ctx}
	req := []student.Course{}
	g.Ctx.ShouldBind(&req)
	
	result := db.MysqlDB.Debug().Create(&req)

	if result.Error != nil {
		g.SendResponse(500, "Fail to add student course", nil)
		return
	}

	g.SendResponse(200,"create student course successfully",nil)

	
}

