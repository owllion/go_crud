package studentRoute

import (
	db "practice/database"
	models "practice/models/student"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)



func UpdateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	
	student := models.Student{}
	
	g.Ctx.ShouldBind(&student)

	result := db.DB.Where(`"ID" = ?`, student.ID).Updates(&student)
	if result.Error != nil {
		g.SendResponse(500,"500",nil)
		return
	}
	g.SendResponse(200, "更新成功", nil)
}