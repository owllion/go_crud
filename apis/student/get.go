package studentRoute

import (
	db "practice/database"
	models "practice/models/student"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)


func GetStudents(ctx *gin.Context) {
	g := handler.GinContext{Ctx:ctx}
	
	students := []models.Student{}
	result := db.DB.Find(&students)

	if result.Error != nil {
		g.SendResponse(500, "500", nil)
		return
	}

	g.SendResponse(200,"獲取成功",students)
}

func GetStudent(ctx *gin.Context) {

	g := handler.GinContext{Ctx: ctx}

	student := models.Student{}

	if id:= ctx.Query("id"); id !="" {
		result := db.DB.Where(`"ID" = ?`, id).Find(&student)

		if result.Error != nil {
			g.SendResponse(500, "500", nil)
			return
		}
		g.SendResponse(200,"200", result)
		return 
	}

	g.SendResponse(200,"獲取成功",student)

}

func SearchStudent(ctx *gin.Context) {
	g := handler.GinContext{Ctx: ctx}

	students  := []models.Student{}
	if keyword := g.Ctx.Query("keyword") ; keyword != "" {
		result :=  db.DB.Where(`"name" = ?"`, keyword).Or(`"email" = ?"`, keyword).Find(&students)

		if result.Error != nil {
			g.SendResponse(500, "500", nil)
			return
		}

		g.SendResponse(200,"ok",result)
	} else {
		g.SendResponse(200,"查詢成功", students)
	}
}