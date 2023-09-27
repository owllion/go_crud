package studentRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)



func DeleteStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	
	student := student.Student{}
	
	g.Ctx.ShouldBind(&student) //populate到struct裡面

	result := db.DB.Debug().Table("student").Delete(`"ID" = ?`, student.ID)

	if result.Error != nil {
		g.SendResponse(500,"刪除失敗",nil)
		return
	}

	g.SendResponse(200,"刪除成功",nil)

}