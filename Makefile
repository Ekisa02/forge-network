BINARY=bin/forge-server

.PHONY: build run test migrate

build:
	go build -o $(BINARY) ./cmd/server

run: build
	./$(BINARY)

test:
	go test ./... -v

migrate:
	# For simple demo we use GORM automigrate at startup. Use a migration tool for production.
	@echo "Migrations are applied at application startup via GORM AutoMigrate."
