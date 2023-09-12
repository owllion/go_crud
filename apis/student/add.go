package studentRoute

import (
	"fmt"
	sql "practice/database"
	models "practice/models/student"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)

func CreateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}

	db := sql.DB
	//創建struct instance
	req := models.Student{}
	
	//從請求拿取資料並populate到空struct裡，type不符會error
	//其他多寫少寫則無視
	g.Ctx.ShouldBind(&req)
	
	fmt.Println(db)

	result := db.Debug().Create(&req)


	if result.Error != nil {
		g.SendResponse(500, "新增失敗", nil)
		return
	}

	g.SendResponse(200,"新增成功",nil)

	
}