package domain

type HealthCheckResponse struct {
	Message string `json:"message"`
}

type HealthCheck struct {
	ID      uint
	Message string
}
