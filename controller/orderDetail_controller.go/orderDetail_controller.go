package orderDetailController

import (
	db "practice/database"
	model "practice/models"
	"practice/util"

	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	*gin.Engine
}

// 構造函數
func NewOrderDetailController(e *gin.Engine) *OrderDetailController {
	return &OrderDetailController{e}
}

func (sc *OrderDetailController) GetCurrencyAndExchangerate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := util.GinContext{Ctx: ctx}

		// var res []model.Order
		var res []model.OrderDetail
		result := db.PostgresDB.Debug().
			// Preload("OrderDetail").
			Where(`orderid = ?`, 2).
			Preload("Order").
			Find(&res)

		if result.Error != nil {
			util.Log("Fail to get orderDetails", "", result.Error.Error())
		}
		if result.RowsAffected == 0 {
			g.SendResponse(400, "No data found", nil)
			return
		}

		g.SendResponse(200, "OK", res)

	}
}
