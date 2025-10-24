package controller

import (
	"net/http"
	"student_management/internal/model"
	"student_management/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service *service.StudentService
}

func NewStudentController(s *service.StudentService) *StudentController {
	return &StudentController{service: s}
}

// create new student
func (c *StudentController) Create(ctx *gin.Context) {
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
