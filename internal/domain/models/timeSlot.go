package models

import "time"

type TimeSlot struct {
	Day   time.Time `json:"day"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
