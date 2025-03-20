package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
	"log"
	"time"
)

type AppointmentService struct {
	calendarRepo repository.WorkerCalendarRepository
	appointRepo  repository.AppointmentRepository
	serviceRepo  repository.ServiceRepository
}

func NewAppointmentService(
	calendarRepo repository.WorkerCalendarRepository,
	appointRepo repository.AppointmentRepository,
	serviceRepo repository.ServiceRepository,
) *AppointmentService {
	return &AppointmentService{
		calendarRepo: calendarRepo,
		appointRepo:  appointRepo,
		serviceRepo:  serviceRepo,
	}
}

func (a *AppointmentService) GetAppointmentSlots(date time.Time, workerID uint, serviceID uint) []models.TimeSlot {
	appointments := a.GetAppointments(date)
	calendar := a.getWorkerCalendar(date, workerID)
	service := a.getService(serviceID, workerID)

	return a.generateSlots(service, appointments, calendar)
}

func (a *AppointmentService) GetAppointments(date time.Time) []models.Appointment {
	appointments, err := a.appointRepo.FindAppointmentsByDay(date)
	if err != nil {
		log.Fatal(err)
	}

	return appointments
}

func (a *AppointmentService) CreateAppointment(dto *dto.AppointmentDto) (*models.Appointment, error) {
	appointment := &models.Appointment{
		WorkerID:  dto.MasterId,
		ClientID:  dto.ClientId,
		ServiceID: dto.ServiceId,
		Date:      dto.Date,
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
		Status:    dto.Status,
	}

	if err := a.appointRepo.Create(appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}

func (a *AppointmentService) getWorkerCalendar(date time.Time, workerID uint) *models.WorkerCalendar {
	workerCalendar, err := a.calendarRepo.FindCalendarByDateAndWorkerID(date, workerID)
	if err != nil {
		log.Fatal(err)
	}

	return workerCalendar
}

func (a *AppointmentService) getService(serviceID uint, workerID uint) *models.Service {
	service, err := a.serviceRepo.FindServiceByIdAndWorker(serviceID, workerID)
	if err != nil {
		log.Fatal(err)
	}

	return service
}

func (a *AppointmentService) generateSlots(
	service *models.Service,
	appointments []models.Appointment,
	workDay *models.WorkerCalendar,
) []models.TimeSlot {
	var slots []models.TimeSlot
	current := workDay.WorkStart //TODO: вынести в константы или структуру, чтобы небыло магических переменных
	step := 30 * time.Minute     //TODO: вынести в константы или структуру, чтобы небыло магических переменных

	for current.Before(workDay.WorkEnd) {

		end := current.Add(service.Duration)

		if end.After(workDay.WorkEnd) {
			break
		}

		isValid := true

		for _, appointment := range appointments {
			if current.Before(appointment.EndTime) && end.After(appointment.StartTime) {
				isValid = false
				break
			}
		}

		if isValid == true {
			slots = append(slots, models.TimeSlot{Day: workDay.Day, Start: current, End: end})
		}

		current = current.Add(step)
	}
	return slots
}
