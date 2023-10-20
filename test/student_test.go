package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	student "practice/models"
	"regexp"

	"net/http/httptest"
	studentRoute "practice/controller/student"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RequestOptions struct {
	Method     string
	Route      string
	Controller gin.HandlerFunc
	Path       string
	// Body       *bytes.Reader
	Body       *bytes.Buffer
}

func doRequest(opts RequestOptions) *httptest.ResponseRecorder {
    r := gin.Default()
    r.Handle(opts.Method, opts.Route, opts.Controller)
	fmt.Println("op.Body", opts.Body)
	if opts.Body == nil {
		opts.Body = &bytes.Buffer{}
	}
    req, _ := http.NewRequest(opts.Method, opts.Path, opts.Body)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    return resp
}



// func getParsedDate(t *testing.T,dateTime string) (parsedDate time.Time, err error) {
// 	const layout = "2006-01-02T15:04:05Z07:00"
// 	parsedDate, err = time.Parse(layout, dateTime)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	return parsedDate, err
// }
func getParsedDate(t *testing.T, dateStr string) (time.Time, error) {
    parsed, err := time.Parse("2006-01-02T15:04:05-07:00", dateStr)
    if err != nil {
        t.Fatalf("Error parsing date: %v", err)
    }
    return parsed.In(time.UTC), nil
}




func TestGetStudent(t *testing.T) {
	SetMock()

	//NOTE: 建立假資料 & sql clause
	parsedBirth, _ := getParsedDate(t,"1995-05-25T08:00:00+08:00")
	fmt.Println("pB",parsedBirth)
	rows := sqlmock.NewRows([]string{"id", "student_id", "name", "birth_date", "admission_year"}).
    AddRow(5, "s123", "Alex", parsedBirth, 2013)

	//NOTE: Postgres Clause
	// query := `SELECT * FROM "enrollment"."student" WHERE "id" = $1`
	// Mock.ExpectQuery(`SELECT (.+) FROM "enrollment"."student" WHERE (.+)`).WithArgs(5).WillReturnRows(rows)


	// queryPattern := regexp.QuoteMeta(`SELECT * FROM "enrollment"."student" WHERE "id" = $1`)
	Mock.ExpectQuery(`SELECT (.+) FROM "enrollment"."student" WHERE (.+)`).WillReturnRows(rows)



	//NOTE: MySQl Clause
	// Mock.ExpectQuery("SELECT * FROM `students` WHERE `id` = ?").WithArgs(5).WillReturnRows(rows)


	getOpts := RequestOptions{
		Method:     "GET",
		Route:      "/student",
		Controller: studentRoute.GetStudent,
		Path:       "/student?id=5",
	}
	resp := doRequest(getOpts)
	

	//NOTE: ExpectationsWereMet 這個一定要放在request之後!!! 不然他就會抱錯說 mock sql clause 沒有match 之類的錯誤
	//NOTE: 加上這，就過不了，不然是有過的
	fmt.Println("RESP!!!", resp)
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

func TestAddStudent(t *testing.T) {
	SetMock()
	
	parsedDate, _ := getParsedDate(t,"1995-05-25T08:00:00+08:00")
	
	query := fmt.Sprintf(`INSERT INTO "enrollment"."student" ("student_id","name","birth_date","admission_year","created_at") VALUES ('s654','John Doe','%s',2019) RETURNING "id","created_at"`, parsedDate.Format("2006-01-02 15:04:05"))


    Mock.ExpectBegin()
	Mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs("s654", "John Doe", parsedDate, 2019)
	Mock.ExpectCommit()
	
		// WillReturnResult(sqlmock.NewResult(1, 1))
	 // assuming the insert will return ID 1 and affect 1 row
	 students := []student.Student{
		{
			StudentID:     "s654",
			Name:          "John Doe",
			BirthDate:     parsedDate,
			AdmissionYear: 2019,
				
		},
	}
	
	jsonData, _ := json.Marshal(students)
	// reader := bytes.NewReader(jsonData)	//這是
	body := bytes.NewBuffer(jsonData)	//這是
	
	postOpts := RequestOptions{
		Method:     "POST",
		Route:      "/student",
		Controller: studentRoute.CreateStudent,
		Path:       "/student",
		Body:       body,
	}
	resp := doRequest(postOpts)
	
	
    // Check if there are any unfulfilled expectations
    if err := Mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }

    // Check the HTTP status code
    expectedStatus := http.StatusOK
    assert.Equal(t, expectedStatus, resp.Code)

    // Compare the expected and actual response
    expectedRes := `{
		"status": 200,
		"msg": "新增成功",
		"data": null
	}`
    var expectedMap, responseMap map[string]interface{}
    json.Unmarshal([]byte(expectedRes), &expectedMap)
    json.Unmarshal([]byte(resp.Body.String()), &responseMap)

	fmt.Println("responseMap:", responseMap)

    assert.Equal(t, expectedMap["msg"], responseMap["msg"])


}
