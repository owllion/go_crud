package apRoute

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestFetchData(t *testing.T) {

// }

// type input struct {
// 	TbLen       int
// 	CurCurrency string
// 	HdCurrency  string
// }

// func TestIsSameCurrency(t *testing.T) {
// 	testCases := []struct {
// 		Name   string
// 		Input  input
// 		Expect bool
// 	}{
// 		{
// 			Name: "Test is same currency",
// 			Input: input{
// 				TbLen:       10,
// 				CurCurrency: "USD",
// 				HdCurrency:  "USD",
// 			},
// 			Expect: true,
// 		},
// 		{
// 			Name: "Test is not the same currency",
// 			Input: input{
// 				TbLen:       10,
// 				CurCurrency: "USD",
// 				HdCurrency:  "IDR",
// 			},
// 			Expect: true,
// 		},
// 		{
// 			Name: "Test is not the same currency",
// 			Input: input{
// 				TbLen:       10,
// 				CurCurrency: "USD",
// 				HdCurrency:  "",
// 			},
// 			Expect: false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(t.Name(), func(t *testing.T) {
// 			actual := isSameCurrency(tc.Input.TbLen, tc.Input.CurCurrency, tc.Input.HdCurrency)
// 			assert.Equal(t, tc.Expect, actual)
// 		})
// 	}
// }
// func TestIsFirstRefered(t *testing.T) {
// 	testCases := []struct {
// 		Name      string
// 		ReqLen    int //要引用的明細list
// 		RecordNum int //當前對帳單的所有明細數量
// 		Expect    bool
// 	}{
// 		{
// 			Name:      "Test is not first reference",
// 			ReqLen:    8,
// 			RecordNum: 3,
// 			Expect:    false,
// 		},
// 		{
// 			Name:      "Test is first reference",
// 			ReqLen:    5,
// 			RecordNum: 0,
// 			Expect:    true,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			actual := isFirstRefered(tc.RecordNum, tc.ReqLen)
// 			assert.Equal(t, tc.Expect, actual)
// 		})
// 	}
// }

// type Input2 struct {
// 	HdOpUUID     string //對帳單ID
// 	CurrencyType string //req傳過來的第一筆明細的幣別
// 	ReqLen       int    //要引用的明細長杜
// 	CurApTbLen   int    //目前對帳單所有明細長度
// }

// // TODO: 去OMS建立銷貨單買幾個料 -> 建一張對帳單並複製他的opUUID給他
// func TestGetCurrencyType(t *testing.T) {
// 	sql.InitialDB()
// 	db := sql.GetWosDB("JU")

// 	testCases := []struct {
// 		Name       string
// 		Input      Input2
// 		ExpectHdCy string
// 	}{
// 		{
// 			Name: "Test is the first reference",
// 			Input: Input2{
// 				HdOpUUID:     "123", //到時候先建立一張空的對帳，複製他的UUID
// 				CurrencyType: "USD", //req第一筆type
// 				ReqLen:       1,
// 				CurApTbLen:   0,
// 			},
// 			ExpectHdCy: "USD",
// 		},
// 		{
// 			Name: "Test is not the first reference",
// 			Input: Input2{
// 				HdOpUUID:     "123", //不用再建一張，因為跑完第一個case就不會是first refer
// 				CurrencyType: "USD", //req第一筆type，要和第一個測試的依樣，
// 				ReqLen:       5,     //不重要 隨便寫
// 				CurApTbLen:   1,     //這其實也不重要，只要不等於0就等於不是第一張
// 			},
// 			ExpectHdCy: "USD",
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			hdCy, errCode, msg := getHdCurrency(db, tc.Input.HdOpUUID, tc.Input.CurApTbLen, tc.Input.CurrencyType, tc.Input.ReqLen)
// 			assert.Equal(t, errCode, 0)
// 			assert.Empty(t, msg)
// 			assert.Equal(t, tc.ExpectHdCy, hdCy) //其實只要不是空的就OK
// 		})
// 	}
// }

// // TODO: 繼上一步驟，創了對帳單後，選好要引用的按下送出，把payload複製，當作reqBody
// func TestTakeOut(t *testing.T) {
// 	//TODO: 複製其他的API TEST
// 	//  /api/takeOut
// 	reqBody := []RequestJSON{
// 		{},
// 		{},
// 	}
// 	//TODO: 其他都依樣，
// }
