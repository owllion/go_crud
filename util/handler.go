package util

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	Ctx *gin.Context
}

type Response struct {
	Status uint16      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

//pointer receiver
func (g *GinContext) SendResponse(status uint16, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Status: status,
		Msg:    msg,
		Data:   data, //map[string]interface{}{},
	})
}
