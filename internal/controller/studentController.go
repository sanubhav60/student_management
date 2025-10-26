package controller

import (
	"net/http"
	"strconv"
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

// get all student
func (c *StudentController) GetAll(ctx *gin.Context) {
	students, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

// get student by id
func (c *StudentController) GetById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := c.service.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, student)
}
