package router

import (
	"fmt"
	studentRoute "practice/apis/student"
	middleware "practice/middleware"

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
		r1.POST("/students", studentRoute.CreateStudent)
		r1.GET("/students/search", studentRoute.SearchStudent)
	}

	return router
}

