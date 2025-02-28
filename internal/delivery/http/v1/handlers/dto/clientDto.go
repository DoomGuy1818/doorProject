package dto

type ClientDto struct {
	FullName string `json:"full_name" validator:"required"`
	Phone    string `json:"phone" validator:"required"`
	Email    string `json:"email" validator:"required"`
	CartId   string `json:"cart_id"`
}
