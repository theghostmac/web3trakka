# Variables
GO_BUILD_CMD=go build -o web3trakka ./cmd/main.go
DOCKER_COMPOSE_CMD=docker-compose -f docker-compose.yml

# Default target
all: build

# Build the application locally
build:
	$(GO_BUILD_CMD)

# Run the application locally
run: build
	./web3trakka

# Build the Docker image
docker-build:
	$(DOCKER_COMPOSE_CMD) build

# Run the application in Docker
docker-up: docker-build
	$(DOCKER_COMPOSE_CMD) up

# Stop the Docker containers
docker-down:
	$(DOCKER_COMPOSE_CMD) down

# Clean up
clean:
	rm -f web3trakka

.PHONY: build run docker-build docker-up docker-down clean