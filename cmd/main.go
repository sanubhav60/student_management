package main

import (
	"student_management/internal/config"
	"student_management/internal/controller"
	"student_management/internal/database"
	"student_management/internal/repository"
	"student_management/internal/router"
	"student_management/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)

	studentRepo := repository.NewStudentRepository(database.DB)
	studentService := service.NewStudentService(studentRepo)
	studentController := controller.NewStudentController(studentService)

	r := router.SetUpRouter(studentController)

	r.Run(":8080")

}
