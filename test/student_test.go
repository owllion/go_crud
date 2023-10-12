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

//NOTE: 理論上大概是這邊
func doRequest(method string, route string, controller gin.HandlerFunc, path string) *httptest.ResponseRecorder {
    r := gin.Default()
    r.Handle(method, route, controller)
    
    req, _ := http.NewRequest(method, path, nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    return resp
}


func getParsedDate(t *testing.T,dateTime string) (parsedDate time.Time, err error) {
	const layout = "2006-01-02T15:04:05Z07:00"
	parsedDate, err = time.Parse(layout, dateTime)
	if err != nil {
		t.Errorf(err.Error())
	}
	return parsedDate, err
}


func TestGetStudent(t *testing.T) {

	//NOTE: 建立假資料 & sql clause
	parsedDate, _ := getParsedDate(t,"1995-05-25T08:00:00+08:00")
	rows := sqlmock.NewRows([]string{"id", "student_id", "name", "birth_date", "admission_year"}).
    AddRow(5, "s123", "Alex", parsedDate, 2013)

	Mock.ExpectQuery("SELECT * FROM `students` WHERE `id` = ?").WithArgs(5).WillReturnRows(rows)

	
	//NOTE:??
	resp := doRequest("GET","/student", studentRoute.GetStudent,"/student?id=5")


	//NOTE: ExpectationsWereMet 這個一定要放在request之後!!! 不然他就會抱錯說 mock sql clause 沒有match 之類的錯誤
	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//NOTE: 確認status code是否為200
	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, resp.Code)
	// fmt.Println("回傳:", resp.Body.String())


	
	//NOTE: 確認回傳資料正確性
	var expectedMap map[string]interface{}
	var responseMap map[string]interface{}

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

	//NOTE:預期資料
	err = json.Unmarshal([]byte(expectedRes), &expectedMap)
	fmt.Println("expectedMap: ", expectedMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal expectedRes: %v", err)
	}

	//NOTE:實際拿到的資料
	err = json.Unmarshal([]byte(resp.Body.String()), &responseMap)
	fmt.Println("responseMap:", responseMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal responseBody: %v", err)
	}

	//NOTE:這邊不能直接用 responseMap["data"][0] 去讀取id，他會說type不對，所以我們才要這樣先cast type，
	data, ok := responseMap["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Failed to assert 'data' into map[string]interface{}")
	}

	//NOTE: 確認型別正確之後才能讀取 key喔!
	sID, exists := data["id"].(float64) 
	// JSON 解碼時，所有數字都會被當作 float64
	
	if !exists {
		t.Fatalf("'id' does not exist in 'data'")
	}

	// 由於 sID 現在是 float64，您可能想將其轉換為整數。
	studentID := int(sID)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, 5, studentID)

	//NOTE:要比較這種傳參考的是否相等，要用一個什麼deepEquals之類的函數，這種一般的Equal不行~
	// assert.Equal(t, expectedMap, responseMap)
}

// func TestAddStudent(t *testing.T) {
	
// 	parsedDate, _ := getParsedDate(t,"1995-05-25T08:00:00+08:00")
// 	rows := sqlmock.NewRows([]string{"id", "student_id", "name", "birth_date", "admission_year"}).
//     AddRow(5, "s123", "Alex", parsedDate, 2013)

// 	Mock.ExpectQuery("SELECT * FROM `students` WHERE `id` = ?").WithArgs(5).WillReturnRows(rows)

	
// 	resp := doRequest("GET","/student", studentRoute.GetStudent,"/student?id=5")


// 	//NOTE: ExpectationsWereMet 這個一定要放在request之後!!! 不然他就會抱錯說 mock sql clause 沒有match 之類的錯誤
// 	if err := Mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}

// 	//NOTE: 確認status code是否為200
// 	expectedStatus := http.StatusOK
// 	assert.Equal(t, expectedStatus, resp.Code)
// 	// fmt.Println("回傳:", resp.Body.String())


	
// 	//NOTE: 確認回傳資料正確性
// 	var expectedMap map[string]interface{}
// 	var responseMap map[string]interface{}

// 	expectedRes := `{
// 		"status": 200,
// 		"msg": "get student data successfully",
// 		"data": {
// 			"id": 5,
// 			"studentId": "s123",
// 			"name": "Alex",
// 			"birthDate": "1995-05-25T08:00:00+08:00",
// 			"admissionYear": 2013,
// 			"courses": null
// 		}
// 	}`

// 	//NOTE:預期資料
// 	err = json.Unmarshal([]byte(expectedRes), &expectedMap)
// 	fmt.Println("expectedMap: ", expectedMap)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal expectedRes: %v", err)
// 	}

// 	//NOTE:實際拿到的資料
// 	err = json.Unmarshal([]byte(resp.Body.String()), &responseMap)
// 	fmt.Println("responseMap:", responseMap)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal responseBody: %v", err)
// 	}

// 	//NOTE:這邊不能直接用 responseMap["data"][0] 去讀取id，他會說type不對，所以我們才要這樣先cast type，
// 	data, ok := responseMap["data"].(map[string]interface{})
// 	if !ok {
// 		t.Fatalf("Failed to assert 'data' into map[string]interface{}")
// 	}

// 	//NOTE: 確認型別正確之後才能讀取 key喔!
// 	sID, exists := data["id"].(float64) 
// 	// JSON 解碼時，所有數字都會被當作 float64
	
// 	if !exists {
// 		t.Fatalf("'id' does not exist in 'data'")
// 	}

// 	// 由於 sID 現在是 float64，您可能想將其轉換為整數。
// 	studentID := int(sID)

// 	assert.Nil(t, err)
// 	assert.True(t, exists)

// }
