package studentRoute

import (
	db "practice/database"
	models "practice/models/student"
	handler "practice/util"

	"github.com/gin-gonic/gin"
)



func DeleteStudent(ctx *gin.Context) {
	//傳id到req.body裡面
	// 還是整個student物件 -> ShouldBind 去populate -> db.Delete(student.ID) 
	g := handler.GinContext{Ctx: ctx}
	
	student := models.Student{}
	
	g.Ctx.ShouldBind(&student) //populate到struct裡面

	result := db.DB.Delete(&student,student.ID)
	if result.Error != nil {
		g.SendResponse(500,"500",nil)
		return
	}

	g.SendResponse(200,"刪除成功",nil)

}