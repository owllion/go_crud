package studentController

import (
	"fmt"

	db "practice/database"
	student "practice/models"
	"practice/util"
	handler "practice/util"
	"reflect"
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
        g := util.GinContext{Ctx: ctx}

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
		g := util.GinContext{Ctx: ctx}
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
		g := util.GinContext{Ctx: ctx}
	
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
		g := util.GinContext{Ctx: ctx}

		type responseJSON struct {
			student.Student
			Name  string `json:"name"`
		}

		// 創建struct instance
		// req := []student.Student{}
		req := []responseJSON{}
		// 從請求拿取資料並populate到空struct裡，type不符會error
		// 其他多寫少寫則無視
		g.Ctx.ShouldBindJSON(&req)
		fmt.Println("this is req", req)

		// 用反射去設置字段為nil，如果它們的值為0或空字符串
		for _, item := range req {
			// 取得item的反射值，因為item是一個結構，所以需要使用Elem()來獲取其基礎值
			// 例如: 如果 item 是 Student{ID: 1, Name: "John"}, 則 v 現在就代表這個Student值
			v := reflect.ValueOf(&item).Elem()
			
			// 迭代這個Student結構的所有欄位
			for i := 0; i < v.NumField(); i++ {
				fieldValue := v.Field(i)  // 這將獲取item的第i個欄位的反射值
				//取得此struct的每個欄位
				
				// 檢查這個欄位是否是指針類型
				if fieldValue.Kind() == reflect.Ptr {
					isZero := false
			
					// 針對欄位的基礎類型檢查其值是否為零值
					switch fieldValue.Elem().Kind() {
					case reflect.String:
						// 如果這個欄位是一個指向字符串的指針，例如Name字段
						// fieldValue.Elem().String() 就會回傳這個字符串的實際值，例如"John"
						isZero = fieldValue.Elem().String() == ""
					case reflect.Float64:
						// 同理，如果欄位是浮點數
						isZero = fieldValue.Elem().Float() == 0
					case reflect.Int, reflect.Int64:
						// 如果是整數，例如 ID，假設它的值是1，這裡就會回傳1
						isZero = fieldValue.Elem().Int() == 0
					case reflect.Struct:
						// 對於結構，我們需要比較它是否等於該類型的零值
						isZero = fieldValue.Elem().Interface() == reflect.Zero(fieldValue.Elem().Type()).Interface()
					}
			
					// 如果欄位是零值，將這個指針設置為nil
					if isZero {
						fieldValue.Set(reflect.Zero(fieldValue.Type()))
					}
				}
			}
		}
		
		result := db.PostgresDB.Debug().Create(&req)

		if result.Error != nil {
			g.SendResponse(500, "新增失敗", nil)
			return
		}

		g.SendResponse(200, "新增成功", nil)
	}

}