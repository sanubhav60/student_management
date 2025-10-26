package dto

type CreateStudentRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateStudentRequest struct {
	Id uint `json:"id" binding:"required"`
}
