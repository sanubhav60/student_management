package service

import (
	"student_management/internal/model"
	"student_management/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) Create(student *model.Student) error {
	return s.repo.Create(student)
}

func (s *StudentService) GetAll() ([]model.Student, error) {
	return s.repo.FindAll()
}
