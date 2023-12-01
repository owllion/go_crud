package aqlController

import (
	db "practice/database"
	student "practice/models"
	"practice/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
type aqlController struct {
    *gin.Engine
}

// 构造函数
func NewAqlController(e *gin.Engine) *aqlController {
    return &aqlController{e}
}

func (sc *aqlController) GetAqlAc() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        g := util.GinContext{Ctx: ctx}
		res := student.Usersamplingplan{}

		var critical_aql_int, major_aql_int, minor_aql_int  float64
		var qty_from_int64, qty_to_int64 int64
		
		inspection_level := g.Ctx.Query("inspection_level")
		critical_aql := g.Ctx.Query("critical_aql")
		major_aql:= g.Ctx.Query("major_aql")
		minor_aql := g.Ctx.Query("minor_aql")

		//NOTE: inspection_level
		// if inspection_level != "" {
		// 	db.PostgresDB = db.PostgresDB.Table(`aql."level_range_letter"."inspection_level" = ?`, inspection_level)
		// }

		//NOTE:  critical_aql
		if critical_aql != "" {
			//轉成數字
			critical_aql_int,_ = strconv.ParseFloat(critical_aql, 64)
			// db.PostgresDB = db.PostgresDB.Table(`aql."aql"."aql" = ?`, critical_aql_int)
		}

		//NOTE:  major_aql
		if major_aql != "" {
			//轉成數字
			major_aql_int,_ = strconv.ParseFloat(major_aql, 64)
			// db.PostgresDB = db.PostgresDB.Table(`aql."aql"."aql" = ?`, major_aql_int)
		}

		//NOTE: minor_aql
		if minor_aql != "" {
			//轉成數字
			minor_aql_int,_ = strconv.ParseFloat(minor_aql,64)
			// db.PostgresDB = db.PostgresDB.Table(`aql."aql"."aql" = ?`, minor_aql_int)
		}

		//NOTE: qty_from
		if qty_from := g.Ctx.Query("qty_from"); qty_from != "" {
			//轉成數字
			qty_from_int64, _ = strconv.ParseInt(qty_from, 10, 64)
			// db.PostgresDB = db.PostgresDB.Table(`aql."product_qty_range"."product_qty_range" = ?`, qty_from_int64)
		}

		//NOTE: qty_to
		if qty_to := g.Ctx.Query("qty_to"); qty_to != "" {
			//轉成數字
			qty_to_int64, _ = strconv.ParseInt(qty_to, 10, 64)

			// db.PostgresDB = db.PostgresDB.Table(`aql."product_qty_range"."qty_to" = ?`, qty_to_int64)
		}

		
		//TODO: 先拿到對應letter(range _to/from + level)
		sample_letter := ""
		result := db.PostgresDB.Debug().
			Table(`aql."level_range_letter"`).
			Select("sampling_letter").
			Where(`"qty_range_from" = ? AND "qty_range_to" = ? AND "inspection_level" = ?`,qty_from_int64, qty_to_int64, inspection_level).
			Find(&sample_letter)
	
		if result.Error != nil {
			g.SendResponse(500, result.Error.Error(), nil)
			return
		}
		if result.RowsAffected == 0 {
			g.SendResponse(404, "No such sampling letter", nil)
			return 
		}

		
		aqls := []float64{
			critical_aql_int,
			major_aql_int,
			minor_aql_int,
		}

		for _, aql := range aqls {

			singlePlan := student.Singleplan{}

			result = db.PostgresDB.Debug().
				Table(`aql."single_plan" as sp`).
				Where(`sp."aql" = ? AND sp."sampling_letter" = ?`, aql, sample_letter).
				Find(&singlePlan)
		
			if result.Error != nil {
				g.SendResponse(500, result.Error.Error(), nil)
				return
			}
			//TODO: 如果沒有找到對應的plan 就會回傳預設值(0)
			// if result.RowsAffected == 0 {
			// 	g.SendResponse(404, "No such sampling plan", nil) 
			// }

			switch aql {
				case critical_aql_int:
					res.Critical_ac = singlePlan.Ac_num
				case major_aql_int:
					res.Major_ac = singlePlan.Ac_num
				case minor_aql_int:
					res.Minor_ac = singlePlan.Ac_num
				default:
					
			}
			

		}


		//TODO: 更新資料到user plan
		res.User_id = 1
		res.Created_at = time.Now().UTC()
		res.Inspection_level = inspection_level
		res.Qty_range_from = qty_from_int64
		res.Qty_range_to = qty_to_int64
		res.Sampling_letter = sample_letter
		res.Critical_aql = critical_aql_int
		res.Major_aql = major_aql_int
		res.Minor_aql = minor_aql_int

		
		result = db.PostgresDB.Debug().
				Table(`aql."user_sampling_plan"`).
				Create(&res)
					
		if result.Error != nil {
			g.SendResponse(500, "Fail to insert new user plan", nil)
			return
		}

		g.SendResponse(200,"get AQL info successfully", res)
		
	}
}
