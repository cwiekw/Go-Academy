build:
	go build -o ./bin/moviemanager_echo/main ./cmd/moviemanager_echo/main.go

run:
	go run ./cmd/moviemanager_echo/main.go

test-unit:
	go test -v ./internal/...

test-k6:
	make docker
	docker run -p 7746:7734 -d --name k6-test go-movie-manager
	k6 run k6-test.js
	docker stop k6-test
	docker rm k6-test

test: test-unit test-k6

docker:
	docker build -t go-movie-manager:latest .

docker-compose:
	docker compose -f docker.yml up -d

docker-compose-down:
	docker compose -f docker.yml down
