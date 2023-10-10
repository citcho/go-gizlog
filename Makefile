.PHONY: help build build-local up down logs ps test migrate rollback generate
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t citcho/gizlog:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with live reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	docker compose exec app go test -v -race -shuffle=on ./...

test-integration:
	docker compose exec app go test -v -tags=integration ./...

initdb: ## Init DB
	docker compose exec app go run ./cmd/migrate/main.go db init

migrate: ## Execute migration
	docker compose exec app go run ./cmd/migrate/main.go db migrate

rollback: ## Execute migration
	docker compose exec app go run ./cmd/migrate/main.go db rollback

generate: ## Generate codes
	docker compose exec app go generate ./...

generate-key: ## Generate key pair for JWT
	openssl genrsa 4096 -out ./internal/common/auth/cert/secret.pem && \
	openssl rsa -pubout -in ./internal/common/auth/cert/secret.pem -out ./internal/common/auth/cert/public.pem

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'