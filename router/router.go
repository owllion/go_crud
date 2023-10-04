package router

import (
	"fmt"
	courseRoute "practice/controller/course"
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

	student := router.Group("api") 
	{
		student.GET("/student", studentRoute.GetStudent)
		student.GET("/students", studentRoute.GetStudents)
		student.POST("/student", studentRoute.CreateStudent)
		student.GET("/students/search", studentRoute.SearchStudent)
		student.DELETE("/student",studentRoute.DeleteStudent)
		student.POST("/student/modify", studentRoute.UpdateStudent)
		student.GET("/search", searchRoute.Search) 
	}

	course := router.Group("api")
	{
		course.GET("/course", courseRoute.GetCourse)
		course.GET("/courses", courseRoute.GetCourses)
		course.DELETE("/course", courseRoute.DeleteCourse)
		course.POST("/course", courseRoute.CreateCourse)
	}


	wsGroup := router.Group("chat")
	{
		wsGroup.GET("/", websocket.HandleConnection)
	}
	return router
}

