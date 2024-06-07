package healthsrv

import (
	"jod/internal/core/domain"
	"jod/internal/core/ports"
	"log"
)

type healthCheckServiceImpl struct {
	repo ports.HealthCheckRepository
}

func NewHealthCheckeService(repo ports.HealthCheckRepository) ports.HealthCheckService {
	return healthCheckServiceImpl{repo: repo}
}

func (s healthCheckServiceImpl) CheckUp() (*domain.HealthCheckResponse, error) {

	happy, err := s.repo.CheckUp()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response := domain.HealthCheckResponse{
		Message: happy.Message,
	}
	return &response, nil
}
