package models

import "time"

type Service struct {
	id       string
	name     string
	isActive bool
	price    float64
	duration time.Time
}

func (s Service) GetId() string {
	return s.id
}

func (s Service) GetName() string {
	return s.name
}

func (s Service) GetPrice() float64 {
	return s.price
}

func (s Service) GetDuration() time.Time {
	return s.duration
}

func (s Service) GetIsActive() bool {
	return s.isActive
}

func (s Service) SetName(name string) {
	s.name = name
}

func (s Service) SetPrice(price float64) {
	s.price = price
}

func (s Service) SetDuration(duration time.Time) {
	s.duration = duration
}
