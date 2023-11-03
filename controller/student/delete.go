package studentController

import (
	db "practice/database"
	student "practice/models"
	"practice/util"

	"strconv"

	"github.com/gin-gonic/gin"
)



func DeleteStudent(ctx *gin.Context) {
	g := util.GinContext{Ctx: ctx}
	
	student := student.Student{}
	
	g.Ctx.ShouldBind(&student)
	if id := ctx.Query("id") ; id != "" {

		id, err := strconv.Atoi(id)

		if err != nil {
			g.SendResponse(400, "無效id", err.Error())
			return
		}
		
		result := db.PostgresDB.Debug().Where("`id` = ?", id).Delete(&student)

		if result.RowsAffected == 0 { 
			g.SendResponse(404, "未找到學生", nil)
			return
		}

		if result.Error != nil {
			g.SendResponse(500,"刪除學生失敗",result.Error.Error())
			return
		}
		g.SendResponse(200, "刪除成功", nil)
		return
	}
	g.SendResponse(400,"請傳學生id",nil)
	
}