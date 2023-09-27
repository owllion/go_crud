package studentRoute

import (
	db "practice/database"
	models "practice/models"
	student "practice/models"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)


func GetStudents(ctx *gin.Context) {
	g := handler.GinContext{Ctx:ctx}
	
	students := []student.Student{}
	result := db.DB.Find(&students)

	if result.Error != nil {
		g.SendResponse(500, "500", nil)
		return
	}

	g.SendResponse(200,"獲取成功",students)
}

func GetStudent(ctx *gin.Context) {

	g := handler.GinContext{Ctx: ctx}

	student := student.Student{}
	if id:= ctx.Query("id"); id !="" {
		result := db.DB.Where(`"ID" = ?`, id).Find(&student)
		
		if result.Error != nil {
			g.SendResponse(500, "500", nil)
			return
		}
		g.SendResponse(200,"200", student)
		return 
	}

	g.SendResponse(200,"回傳所有students",student)

}

func SearchStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}

	students := []models.Student{}
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