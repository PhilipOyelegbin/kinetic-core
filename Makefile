swag:
	swag init -g ./cmd/main.go --parseInternal --parseDependency

run:
	go run ./cmd/main.go

build:
	go build -o ./app

tidy:
	go mod tidy

db-up:
	docker compose up -d

db-down:
	docker compose down --volumes