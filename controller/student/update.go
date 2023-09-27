package studentRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)



func UpdateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	
	student := student.Student{}
	
	g.Ctx.ShouldBind(&student)

	result := db.DB.Debug().Where(`"ID" = ?`, student.ID).Updates(&student)
	//不用寫Select(*)
	
	if result.Error != nil {
		g.SendResponse(500,"更新學生資料失敗",result.Error.Error())
		return
	}
	g.SendResponse(200, "更新成功", nil)
}