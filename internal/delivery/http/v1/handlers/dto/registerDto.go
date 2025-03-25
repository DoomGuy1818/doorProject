package dto

type RegisterDto struct {
	Name           string `json:"name" required:"true"`
	Login          string `json:"login" required:"true"`
	Password       string `json:"password" required:"true"`
	RepeatPassword string `json:"repeat_password" required:"true"`
}
