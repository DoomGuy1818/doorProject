package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/domain/value"
	"doorProject/internal/repository"

	"github.com/labstack/gommon/log"
)

type Client struct {
	repository repository.ClientRepository
}

func NewClientService(repository repository.ClientRepository) *Client {
	return &Client{
		repository: repository,
	}
}

func (c Client) CreateClient(dto *dto.ClientDto) (*models.Client, error) {
	client := &models.Client{
		FullName: dto.FullName,
		Phone:    c.getPhone(dto.Phone),
		Email:    c.getEmail(dto.Email),
	}
	if err := c.repository.Create(client); err != nil {
		return nil, err
	}

	return client, nil
}

func (c Client) getPhone(phone string) value.PhoneNumber {
	phoneNumber, err := value.NewPhoneNumber(phone)
	if err != nil {
		log.Fatal("new phone number error: %s", err.Error())
	}

	return phoneNumber
}

func (c Client) getEmail(email string) value.Email {
	mail, err := value.NewEmail(email)
	if err != nil {
		log.Fatal("new email error: %s", err.Error())
	}

	return mail
}
