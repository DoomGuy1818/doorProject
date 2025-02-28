package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/domain/value"
	"doorProject/internal/repository"
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
		Phone:    value.PhoneNumber(dto.Phone),
		Email:    value.Email(dto.Email),
	}
	if err := c.repository.Create(client); err != nil {
		return nil, err
	}

	return client, nil
}
