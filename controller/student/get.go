package studentController

import (
	db "practice/database"
	student "practice/models"
	"practice/util"
	handler "practice/util"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetStudents(ctx *gin.Context) {
	g := util.GinContext{Ctx:ctx}
	
	students := []student.Student{}
	result := db.PostgresDB.Debug().Find(&students)

	if result.Error != nil {
		g.SendResponse(500, "不明錯誤", result.Error.Error())
		return
	}

	g.SendResponse(200,"get all students data successfully",students)
}

func GetStudent(ctx *gin.Context) {

	g := handler.GinContext{Ctx: ctx}

	student := student.Student{}
	if id:= ctx.Query("id"); id !="" {
		sID, err1 := strconv.Atoi(id)

		if err1 != nil {
			g.SendResponse(400, "Invalid student id", nil)
			return
		}

		//NOTE: MySql寫法
		//其實這邊不把id轉換成int，他也是找的到
		// result := db.PostgresDB.Debug().Where("`id` = ?", sID).Find(&student)

		//NOTE: POstgres寫法
		result := db.PostgresDB.Debug().Where(`"id" = ?`, sID).Find(&student)
		
		if result.Error != nil {
			g.SendResponse(500, result.Error.Error(), nil)
			return
		}
		if result.RowsAffected == 0 {
			g.SendResponse(404, "no such student", nil)
			return 
		}
		g.SendResponse(200,"get student data successfully", student)
		return 
	}

	g.SendResponse(400,"Please provide valid student's id",student)

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