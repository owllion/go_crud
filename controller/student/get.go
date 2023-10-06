package studentRoute

import (
	"fmt"
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)


func GetStudents(ctx *gin.Context) {
	g := handler.GinContext{Ctx:ctx}
	
	students := []student.Student{}
	fmt.Println("db-> ", db.MysqlDB)
	result := db.MysqlDB.Debug().Find(&students)

	if result.Error != nil {
		g.SendResponse(500, "不明錯誤", nil)
		return
	}

	g.SendResponse(200,"獲取成功",students)
}

func GetStudent(ctx *gin.Context) {

	g := handler.GinContext{Ctx: ctx}

	student := student.Student{}
	if id:= ctx.Query("id"); id !="" {
		result := db.MysqlDB.Debug().Where("`id` = ?", id).Find(&student)
		
		if result.Error != nil {
			g.SendResponse(500, "500", nil)
			return
		}
		g.SendResponse(200,"200", student)
		return 
	}

	g.SendResponse(200,"取得所有學生",student)

}

func SearchStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}

	students := []student.Student{}
	if keyword := g.Ctx.Query("keyword") ; keyword != "" {
		result := db.DB.Debug().Where(`name LIKE ?`, "%" + keyword + "%").Or(`email LIKE ?`, "%" + keyword + "%").Find(&students)

		if result.Error != nil {
			g.SendResponse(500, "500", nil)
			return
		}

		g.SendResponse(200,"ok",students)
		return 
	} 

	g.SendResponse(200,"查詢成功", students)
	
}