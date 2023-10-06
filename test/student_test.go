package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	studentRoute "practice/controller/student"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	// 創建 sqlmock 實例
	MysqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	
	defer MysqlDB.Close()

	// 模擬資料庫回應
	rows := sqlmock.NewRows([]string{"id", "student_id", "name", "birth_date", "admission_year"}).
		AddRow(5, "s123", "Alex", "1995-05-25T08:00:00+08:00", 2013)
	mock.ExpectQuery("^SELECT (.+) FROM `students` WHERE `id` = ?$").WithArgs(5).WillReturnRows(rows)

	

	// 使用 gin 測試
	r := gin.Default()
	r.GET("/student", studentRoute.GetStudent)

	req, _ := http.NewRequest("GET", "/student?id=5", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, resp.Code)

	// 使用 ioutil.ReadAll 讀取 response body 的內容
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read resp.Body: %v", err)
	}
	responseBody := string(bodyBytes) // 將讀取到的 bytes 轉換為 string

	fmt.Println("回傳:", responseBody)

	expectedRes := `{
		"status": 200,
		"msg": "200",
		"data": {
			"id": 5,
			"studentId": "S5",
			"name": "Alex",
			"birthDate": "1995-05-25T08:00:00+08:00",
			"admissionYear": 2013,
			"courses": null
		}
	}`

	assert.Equal(t, expectedRes, responseBody)
}
