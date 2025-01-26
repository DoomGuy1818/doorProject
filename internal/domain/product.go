package domain

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
	Color       string
	IsActive    bool
	Size        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
