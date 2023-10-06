package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	studentRoute "practice/controller/student"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func doRequest(method string, route string, controller gin.HandlerFunc, path string) *httptest.ResponseRecorder {
    r := gin.Default()
    r.Handle(method, route, controller)
    
    req, _ := http.NewRequest(method, path, nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    return resp
}


func TestGetStudent(t *testing.T) {
	// // 創建 sqlmock 實例
	// db, mock, err :=  sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) 

	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	
	// defer db.Close()

	

// 在其他代碼上方
const layout = "2006-01-02T15:04:05Z07:00"
parsedDate, err := time.Parse(layout, "1995-05-25T08:00:00+08:00")
if err != nil {
    t.Fatalf("Failed to parse date: %v", err)
}
	// 模擬資料庫回應
	rows := sqlmock.NewRows([]string{"id", "student_id", "name", "birth_date", "admission_year"}).
    AddRow(5, "s123", "Alex", parsedDate, 2013)

	Mock.ExpectQuery("SELECT * FROM `students` WHERE `id` = ?").WithArgs(5).WillReturnRows(rows)


	
	
	//NOTE:原本是　 r = gin.Default() 的 r去設定真的要call的 api
	resp := doRequest("GET","/student", studentRoute.GetStudent,"/student?id=5")

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, resp.Code)

	
	// fmt.Println("回傳:", resp.Body.String())

	expectedRes := `{
		"status": 200,
		"msg": "get student data successfully",
		"data": {
			"id": 5,
			"studentId": "s123",
			"name": "Alex",
			"birthDate": "1995-05-25T08:00:00+08:00",
			"admissionYear": 2013,
			"courses": null
		}
	}`
	
	var expectedMap map[string]interface{}
	var responseMap map[string]interface{}

	err = json.Unmarshal([]byte(expectedRes), &expectedMap)
	fmt.Println("expectedMap: ", expectedMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal expectedRes: %v", err)
	}

	err = json.Unmarshal([]byte(resp.Body.String()), &responseMap)
	fmt.Println("responseMap:", responseMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal responseBody: %v", err)
	}

	data, ok := responseMap["data"].(map[string]interface{})
if !ok {
    t.Fatalf("Failed to assert 'data' into map[string]interface{}")
}

sID, exists := data["id"].(float64) // JSON 解碼時，所有數字都會被當作 float64
if !exists {
    t.Fatalf("'id' does not exist in 'data'")
}

// 由於 sID 現在是 float64，您可能想將其轉換為整數。
studentID := int(sID)

	  // assert if the response are correct
	  assert.Nil(t, err)
	  assert.True(t, exists)
	  assert.Equal(t, 5, studentID)
	  fmt.Println("fsofosfds")

	//不會equal(一班equal不能比較傳參考)
	// assert.Equal(t, expectedMap, responseMap)
}
