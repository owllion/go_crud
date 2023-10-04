package studentRoute

import (
	db "practice/database"
	student "practice/models"
	handler "practice/util"

	"strconv"

	"github.com/gin-gonic/gin"
)



func DeleteCourse(ctx *gin.Context) {

	g := handler.GinContext{Ctx: ctx}
	course := student.Course{}
	g.Ctx.ShouldBind(&course)

	if id := ctx.Query("id") ; id != "" {

		id, err := strconv.Atoi(id)

		if err != nil {
			g.SendResponse(400, "invalid course id", err.Error())
			return
		}
		
		result := db.MysqlDB.Debug().Where("`id` = ?", id).Delete(&course)

		if result.RowsAffected == 0 { 
			g.SendResponse(404, "Course does not exist", nil)
			return
		}

		if result.Error != nil {
			g.SendResponse(500,"Fail to delete the course",result.Error.Error())
			return
		}
		g.SendResponse(200, "Delete the course successfully", nil)
		return
	}
	g.SendResponse(400, "Please provide the student ID.", nil)

	
}