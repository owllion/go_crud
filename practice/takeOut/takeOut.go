package apRoute

// import (
// 	handler "oms_api/controller"
// 	oms_run "oms_api/models/OMS_RUN"
// 	wms_out "oms_api/models/WMS_OUT"
// 	"oms_api/sql"
// 	"oms_api/util"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type RequestJSON struct {
// 	wms_out.Out20rcvhd
// 	HdOpUUID string `json:"HdOpUUID"`
// }

// // 引用應收單
// func TakeOut(ctx *gin.Context) {

// 	g := handler.GinContext{Ctx: ctx}
// 	plantID, _ := g.Ctx.Cookie("plantID")
// 	db := sql.GetWosDB(plantID)
// 	req := []RequestJSON{}

// 	token, _ := g.Ctx.Cookie("token")
// 	updtID := util.DecodeToken(token)

// 	g.Ctx.ShouldBind(&req)

// 	now := time.Now().Local()
// 	sumtotal := 0.0
// 	tax := 0.0
// 	total := 0.0

// 	curApTb := []oms_run.Run81aptb{}
// 	//取得此客戶對帳單明細
// 	errCode, msg := fetchApTbData(db, &curApTb, req[0].HdOpUUID)
// 	if errCode != 0 {
// 		g.SendResponse(errCode, msg, nil)
// 	}

// 	hdCurrencyType, errCode, msg := getHdCurrency(db, req[0].HdOpUUID, len(curApTb), req[0].Currencytype, len(req))
// 	if errCode != 0 {
// 		g.SendResponse(errCode, msg, nil)
// 	}

// 	apTB := []oms_run.Run81aptb{}

// 	//創建新aptb前先檢查 req的每一筆應收的幣別，都要和目前此客戶的某一張對帳單(用hdopUUId去找)的所有aptb皆為相同的
// 	for _, item := range req {
// 		//檢查幣別是否一致
// 		if !isSameCurrency(len(curApTb), item.Currencytype, hdCurrencyType) {
// 			//如果可以進到這邊，代表這張對帳單已寫入幣別，要拿這去判斷(不能用tb，tb沒幣別)
// 			g.SendResponse(400, "幣別不一致，請重新選擇", nil)
// 			return
// 		}

// 		temp := oms_run.Run81aptb{
// 			OpUUID: item.HdOpUUID, // 進料表頭的opUUID
// 			// Itemno:      item.Itemno,
// 			ItemUUID:    util.GenerateUUID(),
// 			PreitemUUID: item.OpUUID,
// 			N14:         item.N14,
// 			N13:         item.N13,
// 			N12:         item.N12,
// 			// N11:        item.N11,
// 			DealN:      item.DealN,
// 			Itemsum:    item.Itemsum,
// 			Itemsumtax: item.Itemsumtax,
// 			Itemtotal:  item.Itemsum + item.Itemsumtax,
// 			// Unitprice:  item.Unitprice,
// 			// Localprice:  item.Localprice,
// 			Hubqty: float64(item.Inqty),
// 			LifeF:  "0",
// 			LifeFT: now,
// 		}
// 		apTB = append(apTB, temp)

// 		sumtotal += item.Itemsum
// 		tax += item.Itemsumtax
// 		total += item.Subtotal

// 		// TODO: 將對帳資料的歷程改為 1
// 		result := db.Debug().Table("2WMS_2OUT.out20rcvhd").
// 			Where(`"opUUID" = ?`, item.OpUUID).
// 			Updates(map[string]interface{}{
// 				"lifeF":  "T",
// 				"lifeFT": now,
// 				"updtT":  now,
// 			})

// 		if result.Error != nil {
// 			util.Log("更新應收單錯誤", item.OpUUID, result.Error.Error())
// 			g.SendResponse(500, "500", nil)
// 			return
// 		}
// 	}

// 	// TODO: 更新對帳單表頭
// 	// 將資料新增至進料表身
// 	if len(apTB) > 0 {
// 		result := db.Table("7OMS_1RUN.run80aphd").
// 			Where(`"opUUID" = ? AND "lifeF" != 'D'`, req[0].HdOpUUID).
// 			Updates(&map[string]interface{}{
// 				"sumtotal": sumtotal,
// 				"tax":      tax,
// 				"total":    total,
// 				"updtID":   updtID,
// 				"updtT":    now,
// 			})

// 		if result.Error != nil {
// 			util.Log("計算金額錯誤", apTB[0].OpUUID, result.Error.Error())
// 			g.SendResponse(500, "500", nil)
// 			return
// 		}

// 		// TODO: 新增對帳單表身
// 		result = db.Table("7OMS_1RUN.run81aptb").Create(&apTB)

// 		if result.Error != nil {
// 			util.Log("引用應收單錯誤", apTB[0].OpUUID, result.Error.Error())
// 			g.SendResponse(500, "500", nil)
// 			return
// 		}
// 	}

// 	g.SendResponse(200, "成功引用應收單", nil)
// }

// func fetchApTbData(db *gorm.DB, curApTB *[]oms_run.Run81aptb, hdOpUUID string) (errCode int, msg string) {
// 	result := db.Debug().Table(`"7OMS_1RUN".run81aptb as run81`).
// 		Where(`run81."opUUID" = ?`, hdOpUUID).
// 		Find(&curApTB)

// 	if result.Error != nil {
// 		return 500, "獲取aptb失敗"
// 	}
// }

// func isFirstRefered(recordNum, reqLen int) bool {
// 	return recordNum == 0 && reqLen > 0
// }

// func updateFirstStatementCurrency(db *gorm.DB, hdOpUUID, currencyType string) (errCode int, msg string) {
// 	result := db.Debug().Table("7OMS_1RUN.run80aphd").
// 		Where(`"opUUID" = ? AND "lifeF" != 'D'`, hdOpUUID).
// 		Updates(&map[string]interface{}{
// 			"currencytype": currencyType,
// 		})

// 	if result.Error != nil {
// 		return 500, "更新對帳單表頭失敗"
// 	}
// }

// func isSameCurrency(tbLen int, curTbCurrencyType, hdCurrencyType string) bool {
// 	if hdCurrencyType == "" {
// 		return false
// 	}
// 	return tbLen > 0 && curTbCurrencyType != hdCurrencyType
// }

// func getHdCurrency(db *gorm.DB, hdOpUUID string, curApTbLen int, currencyType string, reqLen int) (hdCurrencyType string, errCode int, msg string) {
// 	//對帳單創建時無幣別，要等到第一次引用才寫入應收資料的幣別
// 	if isFirstRefered(curApTbLen, reqLen) {
// 		//用第一筆的資料去更新表投幣別
// 		errCode, msg = updateFirstStatementCurrency(db, hdOpUUID, currencyType)

// 		if errCode != 0 {
// 			// g.SendResponse(errCode, msg, nil)
// 			return
// 		}
// 		//更新玩錶投幣別後，就存到這，稍後會用到
// 		hdCurrencyType = currencyType

// 	} else {
// 		//不是第一筆引用，那要自己去DB撈這張對帳的幣別
// 		result := db.Debug().Table(`"7OMS_1RUN".run80aphd`).
// 			Select(`"currencytype"`).
// 			Where(`"opUUID" = ?`, hdOpUUID).
// 			Find(&hdCurrencyType)

// 		if result.Error != nil {
// 			// util.Log("獲取對帳單幣別失敗", nil, result.Error.Error())
// 			// g.SendResponse(500, "獲取對帳單幣別失敗", nil)
// 			errCode = 500
// 			msg = "獲取對帳單幣別失敗"
// 		}
// 	}

// }
