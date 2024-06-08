migrate-local:
	go run ./internal/app/migrate.go

test:
	go test jod/internal/core/...

tidy:
	go mod tidy
	
compose-up:
	docker-compose up --build -d --remove-orphans