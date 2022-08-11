install:
	go mod tidy

run:
	go run main.go

dev:
	air

build:
	go build -o ./bin/main main.go

run-build:
	./bin/main

up:
	docker compose up -d
	docker compose logs -f

upb:
	docker compose up --build -d
	docker compose logs -f

down:
	docker compose down

updb:
	docker-compose -f docker-compose.dev.yml up -d
	docker compose logs -f

downdb:
	docker-compose -f docker-compose.dev.yml down
