package controller

import (
	"net/http"
	"student_management/internal/dto"
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
	var req dto.CreateStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := model.Student{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := c.service.Create(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
