package domain

import "time"

type Order struct {
	Id              string
	Name            string
	IsCanSendNotify bool
	DateStart       time.Time
	DateEnd         time.Time
}
