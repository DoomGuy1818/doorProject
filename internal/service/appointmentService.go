package service

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
	"errors"
	"time"
)

type AppointmentService struct {
	workerRepo   repository.WorkerRepositoryInterface
	calendarRepo repository.WorkerCalendarRepository
	appointRepo  repository.AppointmentRepository
	slotDuration time.Duration
}

func NewAppointmentService(
	workerRepo repository.WorkerRepositoryInterface,
	calendarRepo repository.WorkerCalendarRepository,
	appointRepo repository.AppointmentRepository,
	slotDuration time.Duration,
) *AppointmentService {
	return &AppointmentService{
		workerRepo:   workerRepo,
		calendarRepo: calendarRepo,
		appointRepo:  appointRepo,
		slotDuration: slotDuration,
	}
}

// GetFreeSlots возвращает свободные слоты динамически
func (s *AppointmentService) GetFreeSlots(workerID uint, date time.Time) ([]models.TimeSlot, error) {
	// 1. Получаем рабочие часы
	calendar, err := s.calendarRepo.GetByWorkerAndDate(workerID, date)
	if err != nil {
		return nil, errors.New("рабочий график не найден")
	}

	// 2. Получаем существующие записи
	appointments, err := s.appointRepo.GetByWorkerAndDate(workerID, date)
	if err != nil {
		return nil, errors.New("ошибка получения записей")
	}

	// 3. Генерируем слоты
	return s.generateSlots(calendar, appointments), nil
}

// generateSlots создает слоты на лету
func (s *AppointmentService) generateSlots(
	calendar *models.WorkerCalendar,
	appointments []models.Appointment,
) []models.TimeSlot {
	var slots []models.TimeSlot
	current := time.Date(
		calendar.Day.Year(),
		calendar.Day.Month(),
		calendar.Day.Day(),
		calendar.WorkStart.Hour(),
		calendar.WorkStart.Minute(),
		0, 0, time.UTC,
	)

	endTime := time.Date(
		calendar.Day.Year(),
		calendar.Day.Month(),
		calendar.Day.Day(),
		calendar.WorkEnd.Hour(),
		calendar.WorkEnd.Minute(),
		0, 0, time.UTC,
	)

	for current.Before(endTime) {
		slotEnd := current.Add(s.slotDuration)

		if slotEnd.After(endTime) {
			break
		}

		if s.isSlotAvailable(current, slotEnd, appointments) {
			slots = append(
				slots, models.TimeSlot{
					Start: current,
					End:   slotEnd,
				},
			)
		}

		current = slotEnd
	}

	return slots
}

// Проверка доступности слота
func (s *AppointmentService) isSlotAvailable(
	start, end time.Time,
	appointments []models.Appointment,
) bool {
	for _, app := range appointments {
		if start.Before(app.EndTime) && end.After(app.StartTime) {
			return false
		}
	}
	return true
}
