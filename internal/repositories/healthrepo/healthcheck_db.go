package healthrepo

import (
	"jod/internal/core/domain"
	"jod/internal/core/ports"
)

type healthCheckRepositoryDB struct {
}

func NewHealthCheckRepository() ports.HealthCheckRepository {
	return healthCheckRepositoryDB{}
}

func (r healthCheckRepositoryDB) CheckUp() (*domain.HealthCheck, error) {
	return nil, nil
}
