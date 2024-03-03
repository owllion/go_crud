package studentController

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	gin.SetMode("release")
	r := gin.Default()
	r.GET("/api/hello", NewStudentController(r).GetHello())
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedContent := "Hello"
	//解構response
	var resp struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Fail to parse response json: %v", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedContent, resp.Msg)
}
