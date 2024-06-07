package app

import (
	"jod/internal/core/services/healthsrv"
	"jod/internal/repositories/healthrepo"
)

func Run() {
	healthRopository := healthrepo.NewHealthCheckRepository()
	healthService := healthsrv.NewHealthCheckeService(healthRopository)
	healthService.CheckUp()
}
