package ports

import "jod/internal/core/domain"

type HealthCheckRepository interface {
	CheckUp() (*domain.HealthCheck, error)
}

type HealthCheckService interface {
	CheckUp() (*domain.HealthCheckResponse, error)
}
