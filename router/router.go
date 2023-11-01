package router

import (
	"fmt"
	courseRoute "practice/controller/course"
	orderController "practice/controller/order"
	studentController "practice/controller/student"
	scRoute "practice/controller/student_course"
	middleware "practice/middleware"
	websocket "practice/websocket"

	"github.com/gin-gonic/gin"
)

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.Cors())

	fmt.Println("set up router!")

	studentController := studentController.NewStudentController(router)
	orderController := orderController.NewOrderController(router)

	//NOTE: 這邊 controller 需要"呼叫"getStudent這些函數，因為他們是"回傳"一個gin.HandleFunc，但原本的是"本身" type就是HandleFunc，所以不需要呼叫，直接傳遞即可
	student := router.Group("api") 
	{
		student.GET("/student", studentController.GetStudent())
		student.GET("/students", studentController.GetStudents())
		student.POST("/student", studentController.CreateStudent())
		student.DELETE("/student",studentController.DeleteStudent())
		student.POST("/student/modify", studentController.UpdateStudent())
	}

	order := router.Group("api")
	{
		order.GET("/orders/withinWeek", orderController.GetOrdersWithinAWeek())
		order.GET("/orders/avgPrices", orderController.GetOrderAveragePricesWithinAWeek())
		order.GET("/orders/specificRange", orderController.GetOrderNumInSpecificRange())
		order.GET("/order/countInMonth", orderController.CountOrdersInMonth())
		order.GET("/order/countInEachMonth", orderController.CountOrdersInEachMonth())
	}

	course := router.Group("api")
	{
		course.GET("/course", courseRoute.GetCourse)
		course.GET("/courses", courseRoute.GetCourses)
		course.DELETE("/course", courseRoute.DeleteCourse)
		course.POST("/course", courseRoute.CreateCourse)
	}

	studentCourseRoute := router.Group("api")
	{
		studentCourseRoute.GET("/sc", scRoute.GetStudentCourses)
		studentCourseRoute.POST("/sc", scRoute.CreateStudentCourse)
		studentCourseRoute.DELETE("/sc", scRoute.DeleteStudentCourse)
	}


	wsGroup := router.Group("chat")
	{
		wsGroup.GET("/", websocket.HandleConnection)
	}
	return router
}

