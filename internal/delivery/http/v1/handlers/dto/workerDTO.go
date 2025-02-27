package dto

type WorkerDTO struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=8"`
}
