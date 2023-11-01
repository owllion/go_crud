package orderController

import (
	"fmt"
	"math/rand"
	db "practice/database"
	student "practice/models"
	handler "practice/util"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
type OrderController struct {
    *gin.Engine
}

// 構造函數
func NewOrderController(e *gin.Engine) *OrderController {
    return &OrderController{e}
}


// 商業邏輯

func (sc *OrderController) CountOrdersInEachMonth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type responseJSON struct {
			TotalNumOfOrders map[string]int
		}
		g := handler.GinContext{Ctx: ctx}
		res := responseJSON{}

		tempTotalNumOfOrders := make(map[string]int)

		type TotalNumRes struct {
			TotalNum int `gorm:"column:total_num_of_orders"`
		}

		eachTotalNum := TotalNumRes{}
		//迴圈裡面建每個月初~底
		for i:=1; i<13; i++ {
			startDate := time.Date(time.Now().Year(), time.Month(i), 1, 0, 0, 0, 0, time.UTC)
			endDate := startDate.AddDate(0, 1, -1)
			startDateStr := startDate.Format("2006-01-02") + " 00:00:00"
			endDateStr := endDate.Format("2006-01-02") + " 23:59:59"


			result := db.PostgresDB.
					Debug().
					Table(`enrollment."order"`).
					Select(`count(*) OVER () as total_num_of_orders`).
					Where(`"created_at" >= ? AND "created_at" <= ?`, startDateStr, endDateStr).
					Order(`"created_at"`).
					Find(&eachTotalNum)

			if result.Error != nil {
				g.SendResponse(500, "Fail to get data", result.Error.Error())
				return
			}

			// if result.RowsAffected == 0 {
			// 	g.SendResponse(400, "No data found", nil)
			// 	return 
			// }

			tempTotalNumOfOrders[strconv.Itoa(i)] = eachTotalNum.TotalNum
			fmt.Println("---------------ea.Total-----------------", eachTotalNum.TotalNum)

		}
		res.TotalNumOfOrders = tempTotalNumOfOrders

		g.SendResponse(200, "OK!", res)
	}
}
func (sc *OrderController) CountOrdersInMonth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type responseJSON struct {
			TotalNumOfOrders int `gorm:"column:total_number_of_orders"`
		}
		
		g := handler.GinContext{Ctx: ctx}
		res := responseJSON{} 

		var startDate,endDate time.Time
		var startDateStr, endDateStr string
		
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(11)

		if randomNumber < 8 {
			//有傳
			monthStr := "5"
			month, err := strconv.Atoi(monthStr)
			if err != nil {
				g.SendResponse(400,"Fail to convert month string", err.Error())
				return
			}
			//月初
			startDate = time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			
			startDateStr = startDate.Format("2006-01-02")

			//月底
			endDate = startDate.AddDate(0,1,-1)
			endDateStr = endDate.Format("2006-01-02")
		}
		
		//前端傳的月份，沒傳就是當月
		result := db.PostgresDB.
					Debug().
					Table(`enrollment."order"`).
					Select(`count(*) as total_number_of_orders`).
					Where(`"created_at" >= ? AND "created_at" <= ?`, startDateStr, endDateStr).
					Find(&res)

		if result.Error != nil {
			g.SendResponse(500,"Fail to retrieve data",result.Error.Error())
			return
		}
		if result.RowsAffected == 0 {
			g.SendResponse(400, "No data found", nil)
			return
		}
		g.SendResponse(200, "Get data successfully", res)
		
	}
}
func (sc *OrderController) GetOrderNumInSpecificRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO: 找出特定範圍，例如2023-01-01到2023-03-31，每天的訂單數量
		g := handler.GinContext{Ctx : ctx}
		res := []student.Order{}
	
		var startDate,endDate string

		// 生成一个范围在0到10之间的随机整数
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(11)

		if randomNumber > 5 {
			//模擬前端傳日期(就只有日期，沒有時間，透過Query)
			startDate = "2023-01-01"
			endDate = "2023-03-31"
		}else {
			//沒傳就預設這個月初~月底
			currentTime := time.Now()
			firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1,0,0,0,0,currentTime.Location())
			lastDayOfMonth := firstDayOfMonth.AddDate(0,1,-1) //這個月最後一天
			/*
			NOTE:
				函數:
					func (t Time) AddDate(years int, months int, days int) Time
					- years 参数表示要添加的年数。
					- months 参数表示要添加的月数。
					- days 参数表示要添加的天数。
				-------------------------------------------------------------
				使用:
					0 表示不添加年份，保持年份不變。
					1 表示添加1個月。
					-1 表示再减去1天，以獲取當前月份的最後一天。
				-------------------------------------------------------------
				这是一种常见的方法来计算当前月份的最后一天，因为不同月份的天数不同，所以不能简单地假设每个月都是30或31天。
			*/

			startDate = firstDayOfMonth.Format("2006-01-02")
			endDate = lastDayOfMonth.Format("2006-01-02")

			fmt.Println("sDate ~ eDate", startDate, endDate)
		}

		
		result := db.PostgresDB.Debug().
					Table(`enrollment."order"`).
					Where(`"created_at" >= ? AND "created_at" <= ?`, startDate + " 00:00:00", endDate + " 23:59:59").
					Order(`"created_at"`).
					Find(&res)

		if result.Error != nil {
			g.SendResponse(500, "Fail to get data", result.Error.Error())
			return
		}

		if result.RowsAffected == 0 {
			g.SendResponse(400,"Data not found", nil)
			return 
		}

	g.SendResponse(200, "Successfully get the data", res)	
		
	}
}
func (sc *OrderController) GetOrderAveragePricesWithinAWeek() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        g := handler.GinContext{Ctx: ctx}

		//TODO: 取得7天內每一天的total_price平均值(formatted to one decimal place)，並回傳map(k:日期, v:平均值)
		queryType := "average prices"
		// avgPrices := make(map[string]interface{})
		today := time.Now()
		sixDaysAgo := today.AddDate(0,0,-6) 
		
		type responseJSON struct {
			AvgPrices map[string]float64 `json:"avgPrices"`
		}
		response := responseJSON{
			AvgPrices: make(map[string]float64),
		} 
		/*
		NOTE: 
		上面那樣初始化，裡面所有field都會是他的初始值，map那種就會直接是nil，所以如果直接response.AvgPrices[xx] = oooo 就會報錯說你賦值給nil map。
		解法
			1.另寫一個map來暫存拿到/算出的值，最後跳出迴圈時再賦完整map給response就不會error
			2.手動初始化 ->
				response := responseJSON{
    				AvgPrices: make(map[string]interface{}),
				}
		*/
		fmt.Println("--------response.AvgPrices", response.AvgPrices)//nil
			
		// 傻爆眼 = = 搞了我快兩天 結果是因為db欄位名稱(還是我自訂的咧)和接收的struct欄位名稱不符合 = =  又是gorm tag問題:)) 真棒欸
		type Result struct {
			AveragePrice float64 `gorm:"column:average_price"`
		}

		avgResult := Result{}
		for i := 0; i < 7; i++ {
			//從該天00:00:00開始，到 23:59:59結束
			// startDate = time.Date(today.Year(), today.Month(), today.Day()-i, 0,0,0,0,time.UTC)
			// endDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
			
			//7天前
			//NOTE: 時間用format較好
			currentDate := sixDaysAgo.AddDate(0,0,i) 

			formattedDate := currentDate.Format("2006-01-02")

			startDate := formattedDate + " 00:00:00"
			endDate := formattedDate + " 23:59:59"
			
			// avgMap := make(map[string]float64)

			//只找單獨一天的price平均
			result :=  db.PostgresDB.Debug().
				Table(`enrollment."order"`).
				Select(`ROUND(AVG("total_price")::numeric,1) as average_price`).
				Where(`"created_at" >= ? AND "created_at" <= ?`, startDate, endDate).
				Find(&avgResult)
		
			
			fmt.Println("-------------db result",avgResult)
			
			// if value, ok := avgMap["average_price"]; ok {
			// 	fmt.Println("印出type--------", reflect.TypeOf(value)) //float64 = = 
			// 	avgPrices[formattedDate] = value
			// } else {
			// 	fmt.Println("average_price not found in resultMap")
			// }
			
		
			// if val, ok := avgMap["average_price"]; ok {
			// 	// floatPrice, err := strconv.ParseFloat(val.(string), 64)

			// 	// if err != nil {
			// 	// 	fmt.Println("轉換錯誤")
			// 	// 	return 
			// 	// }
			// 	response.AvgPrices[formattedDate] = val
			// 	fmt.Println("response.AvgPrices", response.AvgPrices)

			// }
			
			response.AvgPrices[formattedDate] = avgResult.AveragePrice

			if result.Error != nil {
				g.SendResponse(500, result.Error.Error(), nil)
				return
			}
			if result.RowsAffected == 0 {
				g.SendResponse(404, "no such orders", nil)
				return 
			}
			
		}
		
		// response.AvgPrices = avgPrices
		
		// if result.Error != nil {
		// 	g.SendResponse(500, result.Error.Error(), nil)
		// 	return
		// }
		// if result.RowsAffected == 0 {
		// 	g.SendResponse(404, "no such orders", nil)
		// 	return 
		// }
		g.SendResponse(200,fmt.Sprintf("get %s data successfully", queryType), response)
		 
	}

}
func (sc *OrderController) GetOrdersWithinAWeek() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        g := handler.GinContext{Ctx: ctx}
		res := []student.Order{}

		//TODO:選出最近一週的所有訂單。
		queryType := "recent week"
		//前7天 00:00:00
		startDate := time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day()-7,0,0,0,0,time.UTC)
		//今天 23:59:59
		endDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.UTC)

		result := db.PostgresDB.Debug().
			Where(`"created_at" >= ? AND "created_at" <= ?`, startDate, endDate).
			Order("created_at, total_price").
			Find(&res)

		

			if result.Error != nil {
				g.SendResponse(500, result.Error.Error(), nil)
				return
			}
			if result.RowsAffected == 0 {
				g.SendResponse(404, "no such orders", nil)
				return 
			}
			
		g.SendResponse(200,fmt.Sprintf("get %s data successfully", queryType), res)
		 
	}

}

func (sc *OrderController) GetOrder() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        g := handler.GinContext{Ctx: ctx}

		Order := student.Order{}
		if id:= ctx.Query("id"); id !="" {
			sID, err := strconv.Atoi(id) //轉成int

			if err != nil {
				g.SendResponse(400, "Invalid Order id", nil)
				return
			}

			//NOTE: MySql寫法
			//其實這邊不把id轉換成int，他也是找的到
			// result := db.PostgresDB.Debug().Where("`id` = ?", sID).Find(&Order)

			//NOTE: POstgres寫法
			result := db.PostgresDB.Debug().Where(`"id" = ?`, sID).Find(&Order)
			
			if result.Error != nil {
				g.SendResponse(500, result.Error.Error(), nil)
				return
			}
			if result.RowsAffected == 0 {
				g.SendResponse(404, "no such Order", nil)
				return 
			}
			g.SendResponse(200,"get Order data successfully", Order)
			return 
		}

		g.SendResponse(400,"Please provide valid Order's id",Order)
	}
}

func (sc *OrderController) GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := handler.GinContext{Ctx:ctx}
		Orders := []student.Order{}
		result := db.PostgresDB.Debug().Find(&Orders)

		if result.Error != nil {
			g.SendResponse(500, "不明錯誤", result.Error.Error())
			return
		}

		g.SendResponse(200,"get all Orders data successfully",Orders)

	}
}

func (sc *OrderController) DeleteOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := handler.GinContext{Ctx: ctx}
		Order := student.Order{}

		g.Ctx.ShouldBind(&Order)

		if id := ctx.Query("id") ; id != "" {

			id, err := strconv.Atoi(id)

			if err != nil {
				g.SendResponse(400, "無效id", err.Error())
				return
			}
			
			result := db.PostgresDB.Debug().Where("`id` = ?", id).Delete(&Order)

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


func (sc *OrderController) UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context)  {
		g := handler.GinContext{Ctx: ctx}
	
		Order := student.Order{}
		
		g.Ctx.ShouldBind(&Order)
		if id := ctx.Query("id") ; id != "" {
	
			id, err := strconv.Atoi(id)
	
			if err != nil {
				g.SendResponse(400, "無效id", err.Error())
				return
			}
			//NOTE: mysql的寫法是"‵‵"，剛好和postgres相反(大概啦...目前改完後就更新成功了)
			result := db.PostgresDB.Debug().Where("`id` = ?", id).Updates(&Order)
	
			if result.RowsAffected == 0 {  //沒找到資料不會抱錯
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


func(sc * OrderController) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := handler.GinContext{Ctx: ctx}

		type responseJSON struct {
			student.Order
			Name  string `json:"name"`
		}

		// 創建struct instance
		// req := []Order.Order{}
		req := []responseJSON{}

		g.Ctx.ShouldBindJSON(&req)
		fmt.Println("this is req", req)

		// 用反射去設置字段為nil，如果它們的值為0或空字符串
		for _, item := range req {
			// 取得item的反射值，因為item是一個結構，所以需要使用Elem()來獲取其基礎值
			// 例如: 如果 item 是 Order{ID: 1, Name: "John"}, 則 v 現在就代表這個Order值
			v := reflect.ValueOf(&item).Elem()
			
			// 迭代這個Order結構的所有欄位
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