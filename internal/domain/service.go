package domain

import "time"

type Service struct {
	Id       string
	Name     string
	isActive bool
	price    float64
	duration time.Time
}
