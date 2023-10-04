package router

import (
	"fmt"
	searchRoute "practice/controller/search"
	studentRoute "practice/controller/student"
	middleware "practice/middleware"
	websocket "practice/websocket"

	"github.com/gin-gonic/gin"
)

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.Cors())

	fmt.Println("set up router!")

	r1 := router.Group("api") 
	{
		r1.GET("/student", studentRoute.GetStudent)
		r1.GET("/students", studentRoute.GetStudents)
		r1.POST("/student", studentRoute.CreateStudent)
		r1.GET("/students/search", studentRoute.SearchStudent)
		r1.DELETE("/student",studentRoute.DeleteStudent)
		r1.POST("/student/modify", studentRoute.UpdateStudent)
		r1.GET("/search", searchRoute.Search) 
	}
	wsGroup := router.Group("chat")
	{
		wsGroup.GET("/", websocket.HandleConnection)
	}
	return router
}

