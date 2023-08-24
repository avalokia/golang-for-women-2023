package routers

import (
	"assignment-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()
	router.GET("/students", controllers.GetAllStudents)
	router.POST("/student", controllers.CreateStudent)
	router.PUT("/student/:id", controllers.UpdateStudentByID)
	router.DELETE("/student/:id", controllers.DeleteStudentByID)

	return router
}
