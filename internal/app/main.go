package app

import (
	"jod/internal/adapters/repositories/healthrepo"
	"jod/internal/core/services/healthsrv"
)

func Run() {
	healthRopository := healthrepo.NewHealthCheckRepository()
	healthService := healthsrv.NewHealthCheckeService(healthRopository)
	healthService.CheckUp()
}
