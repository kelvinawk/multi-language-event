include .env
MIGRATIONS_PATH = ./db/migration

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_DSN) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_DSN) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: start
start:
	@go run ./cmd/api/
