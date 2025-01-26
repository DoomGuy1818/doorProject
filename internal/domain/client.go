package domain

type Client struct {
	Id       string
	FullName string
	Phone    string //TODO: вынести номер телефона в Value-object
	Email    string //TODO: вынести мыло в Value-object
}
