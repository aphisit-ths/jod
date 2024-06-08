package healthsrv_test

import (
	"errors"
	"jod/internal/adapters/repositories/healthrepo"
	"jod/internal/core/domain"
	"jod/internal/core/services/healthsrv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanHealthCheckCorrectly(t *testing.T) {
	expectedMessage := "I'm fine"
	expectedObject := domain.HealthCheck{ID: 1, Message: expectedMessage}

	mockRepo := healthrepo.NewHealthCheckRepositoryMock()
	mockRepo.On("CheckUp").Return(&expectedObject, nil)

	service := healthsrv.NewHealthCheckeService(mockRepo)

	response, _ := service.CheckUp()
	assert.Equal(t, expectedMessage, response.Message)
}
func TestCanHealthCheckWithError(t *testing.T) {
	expectedObject := domain.HealthCheck{}
	expectedErrorMessage := "unhappy health check !"
	mockRepo := healthrepo.NewHealthCheckRepositoryMock()
	mockRepo.On("CheckUp").Return(&expectedObject, errors.New(expectedErrorMessage))

	healthCheckService := healthsrv.NewHealthCheckeService(mockRepo)

	response, err := healthCheckService.CheckUp()

	assert.Errorf(t, err, expectedErrorMessage)
	assert.Nil(t, response)
}
