package studentRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"
	"strconv"

	"github.com/gin-gonic/gin"
)



func UpdateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	
	student := student.Student{}
	
	g.Ctx.ShouldBind(&student)
	if id := ctx.Query("id") ; id != "" {

		id, err := strconv.Atoi(id)

		if err != nil {
			g.SendResponse(400, "無效id", err.Error())
			return
		}
		//NOTE: mysql的寫法是"‵‵"，剛好和postgres相反(大概啦...目前改完後就更新成功了)
		result := db.PostgresDB.Debug().Where("`id` = ?", id).Updates(&student)

		if result.RowsAffected == 0 { 
			/*
				NOTE: 
					用gorm的updates，如果找不到資源，他就只是不更新，不會抱錯，還會return 200! 如果想要確切知道到底有沒有資料被更新，就要用RowsAffected去查看被影響的資料筆數
			*/

			g.SendResponse(404, "未找到學生", nil)
			return
		}

		if result.Error != nil {
			g.SendResponse(500,"更新學生資料失敗",result.Error.Error())
			return
		}
		g.SendResponse(200, "更新成功", nil)
		return
	}
	g.SendResponse(400,"請傳學生id",nil)
	
}