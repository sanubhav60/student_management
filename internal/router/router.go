package router

import (
	"student_management/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(studentController *controller.StudentController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/student", studentController.Create)
	}

	return r
}
