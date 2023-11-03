package courseRoute

import (
	db "practice/database"
	student "practice/models"
	"practice/util"

	"github.com/gin-gonic/gin"
)


func GetCourses(ctx *gin.Context) {
	g := util.GinContext{Ctx: ctx}	
	courses := []student.Course{}
	result := db.MysqlDB.Debug().Find(&courses)

	if result.Error != nil {
		g.SendResponse(500, "Fail to get all the courses data", nil)
		return
	}

	g.SendResponse(200,"get all courses",courses)
}

func GetCourse(ctx *gin.Context) {

	g := util.GinContext{Ctx: ctx}
	course := student.Course{}
	if id:= ctx.Query("id"); id !="" {
		result := db.DB.Where(`"ID" = ?`, id).Find(&course)
		
		if result.Error != nil {
			g.SendResponse(500, "fail to get the course", nil)
			return
		}
		g.SendResponse(200,"get the course", course)
		return 
	}

	g.SendResponse(400,"Please pass the course id", nil)

}