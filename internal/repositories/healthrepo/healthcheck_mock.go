package healthrepo

import (
	"jod/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type healthCheckRepositoryMock struct {
	mock.Mock
}

func NewHealthCheckRepositoryMock() *healthCheckRepositoryMock {
	return &healthCheckRepositoryMock{}
}

func (m healthCheckRepositoryMock) CheckUp() (*domain.HealthCheck, error) {
	args := m.Called()
	return args.Get(0).(*domain.HealthCheck), args.Error(1)
}
