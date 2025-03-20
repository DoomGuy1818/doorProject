package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type WorkerCalendar struct {
	repository repository.WorkerCalendarRepository
}

func NewWorkerCalendar(repo repository.WorkerCalendarRepository) *WorkerCalendar {
	return &WorkerCalendar{repository: repo}
}

func (wc *WorkerCalendar) CreateWorkDay(dto *dto.WorkerCalendarDto) (*models.WorkerCalendar, error) {

	workDay := &models.WorkerCalendar{
		Day:       dto.Day,
		WorkStart: dto.WorkStart,
		WorkEnd:   dto.WorkEnd,
		WorkerID:  dto.WorkerId,
	}

	if err := wc.repository.Create(workDay); err != nil {
		return nil, err
	}

	return workDay, nil
}
