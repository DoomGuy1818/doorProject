package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type WorkerService struct {
	workerRepository repository.WorkerRepositoryInterface
}

func NewWorkerService(workerRepository repository.WorkerRepositoryInterface) *WorkerService {
	return &WorkerService{
		workerRepository: workerRepository,
	}
}

func (w WorkerService) CreateWorker(dto *dto.WorkerDTO) (*models.Worker, error) {
	passwordHash := generatePasswordHash(dto.Password)

	err := comparePasswordHash(passwordHash, dto.RepeatPassword)
	if err != nil {
		log.Fatal(err.Error())
	} //TODO: вместо фатала, возвращать nil, err, чтобы было что выводит хендлер

	worker := models.Worker{
		Name:     dto.Name,
		Password: dto.Password,
		Login:    dto.Email,
	}

	if err = w.workerRepository.Create(&worker); err != nil {
		log.Fatal(err.Error())
	}
	return &worker, err
}

func generatePasswordHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(bytes)
}

func comparePasswordHash(password string, repeatPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(repeatPassword))
}
