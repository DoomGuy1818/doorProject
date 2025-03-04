package models

import "time"

type TimeSlot struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
