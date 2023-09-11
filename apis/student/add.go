package studentRoute

import (
	db "practice/database"
	models "practice/models/student"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)

func CreateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}
	
	//創建struct instance
	req := models.Student{}
	
	//從請求拿取資料並populate到空struct裡，type不符會error
	//其他多寫少寫則無視
	g.Ctx.ShouldBind(&req)
	
	result := db.DB.Create(&req)

	if result.Error != nil {
		g.SendResponse(500, "500", nil)
		return
	}

	g.SendResponse(201,"新增成功",nil)

	
}