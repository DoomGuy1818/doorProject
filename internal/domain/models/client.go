package models

import "doorProject/internal/domain/value"

type Client struct {
	id       string
	fullName string
	phone    value.PhoneNumber
	email    value.Email
}

func (c Client) GetId() string {
	return c.id
}

func (c Client) GetName() string {
	return c.fullName
}

func (c Client) GetPhone() value.PhoneNumber {
	return c.phone
}

func (c Client) GetEmail() value.Email {
	return c.email
}

func (c Client) SetName(name string) {
	c.fullName = name
}

func (c Client) SetPhone(phone value.PhoneNumber) {
	c.phone = phone
}

func (c Client) SetEmail(email value.Email) {
	c.email = email
}
