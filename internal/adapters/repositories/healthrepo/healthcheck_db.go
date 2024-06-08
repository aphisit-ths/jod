package healthrepo

import (
	"errors"
	"jod/internal/core/domain"
	"jod/internal/core/ports"
)

type healthCheckRepository struct {
}

func NewHealthCheckRepository() ports.HealthCheckRepository {
	return healthCheckRepository{}
}

func (r healthCheckRepository) CheckUp() (*domain.HealthCheck, error) {
	return nil, errors.New("not Implemented")
}
