package studentRoute

import (
	"fmt"
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)

func CreateStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}

	//創建struct instance
	req := []student.Student{}
	fmt.Println("this is req",req)
	
	//從請求拿取資料並populate到空struct裡，type不符會error
	//其他多寫少寫則無視
	g.Ctx.ShouldBindJSON(&req)	
	result := db.PostgresDB.Debug().Create(&req)


	if result.Error != nil {
		g.SendResponse(500, "新增失敗", nil)
		return
	}

	g.SendResponse(200,"新增成功",nil)

	
}

