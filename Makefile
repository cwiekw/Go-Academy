build:
	go build -o ./bin/moviemanager_echo/main ./cmd/moviemanager_echo/main.go

run:
	go run ./cmd/moviemanager_echo/main.go

test:
	go test -v ./internal/...

docker:
	docker build -t go-movie-manager:latest .

docker-compose:
	docker compose -f docker.yml up -d

