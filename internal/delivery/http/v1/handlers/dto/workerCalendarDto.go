package dto

import "time"

type WorkerCalendarDto struct {
	Day       time.Time `json:"day" validate:"required"`
	WorkStart time.Time `json:"work_start" validate:"required"`
	WorkEnd   time.Time `json:"work_end" validate:"required"`
	WorkerId  uint      `json:"worker_id" validate:"required"`
}
