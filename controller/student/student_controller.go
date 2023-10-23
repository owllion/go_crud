package studentController

import (
	"fmt"
	db "practice/database"
	student "practice/models"
	handler "practice/util"
	"strconv"

	"github.com/gin-gonic/gin"
)
type StudentController struct {
    *gin.Engine
}

// 构造函数
func NewStudentController(e *gin.Engine) *StudentController {
    return &StudentController{e}
}


// 这里是业务方法
func (sc *StudentController) GetStudent() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        g := handler.GinContext{Ctx: ctx}

		student := student.Student{}
		if id:= ctx.Query("id"); id !="" {
			sID, err := strconv.Atoi(id) //轉成int

			if err != nil {
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
}


func (sc *StudentController) GetStudents() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := handler.GinContext{Ctx:ctx}
		students := []student.Student{}
		result := db.PostgresDB.Debug().Find(&students)

		if result.Error != nil {
			g.SendResponse(500, "不明錯誤", result.Error.Error())
			return
		}

		g.SendResponse(200,"get all students data successfully",students)

	}
}

func (sc *StudentController) DeleteStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := handler.GinContext{Ctx: ctx}
		student := student.Student{}

		g.Ctx.ShouldBind(&student)

		if id := ctx.Query("id") ; id != "" {

			id, err := strconv.Atoi(id)

			if err != nil {
				g.SendResponse(400, "無效id", err.Error())
				return
			}
			
			result := db.PostgresDB.Debug().Where("`id` = ?", id).Delete(&student)

			if result.RowsAffected == 0 { 
				g.SendResponse(404, "未找到學生", nil)
				return
			}

			if result.Error != nil {
				g.SendResponse(500,"刪除學生失敗",result.Error.Error())
				return
			}
			g.SendResponse(200, "刪除成功", nil)
			return
		}
		g.SendResponse(400,"請傳學生id",nil)
	}
}


func (sc *StudentController) UpdateStudent() gin.HandlerFunc {
	return func(ctx *gin.Context)  {
		g := handler.GinContext{Ctx: ctx}
	
		student := student.Student{}
		
		g.Ctx.ShouldBind(&student)
		if id := ctx.Query("id") ; id != "" {
	
			id, err := strconv.Atoi(id)
	
			if err != nil {
				g.SendResponse(400, "無效id", err.Error())
				return
			}
			//NOTE: mysql的寫法是"‵‵"，剛好和postgres相反(大概啦...目前改完後就更新成功了)
			result := db.PostgresDB.Debug().Where("`id` = ?", id).Updates(&student)
	
			if result.RowsAffected == 0 { 
				/*
				NOTE: 
					用gorm的updates，如果找不到資源，他就只是不更新，不會抱錯，還會return 200! 如果想要確切知道到底有沒有資料被更新，就要用RowsAffected去查看被影響的資料筆數
				*/
	
				g.SendResponse(404, "未找到學生", nil)
				return
			}
	
			if result.Error != nil {
				g.SendResponse(500,"更新學生資料失敗",result.Error.Error())
				return
			}
			g.SendResponse(200, "更新成功", nil)
			return
		}
		g.SendResponse(400,"請傳學生id",nil)

	}
}


func(sc * StudentController) CreateStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
}