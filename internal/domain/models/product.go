package models

import "time"

type Product struct {
	id          string
	name        string
	description string
	price       float64
	color       string
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (p Product) GetId() string {
	return p.id
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetDescription() string {
	return p.description
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetColor() string {
	return p.color
}

func (p Product) GetIsActive() bool {
	return p.isActive
}

func (p Product) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p Product) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p Product) SetName(name string) {
	p.name = name
}

func (p Product) SetDescription(description string) {
	p.description = description
}

func (p Product) SetPrice(price float64) {
	p.price = price
}

func (p Product) SetColor(color string) {
	p.color = color
}

func (p Product) SetIsActive(isActive bool) {
	p.isActive = isActive
}

func (p Product) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p Product) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}
