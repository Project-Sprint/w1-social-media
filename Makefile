.PHONY: migrate-up migrate-down

run:
	go run cmd/main.go

migrate-up:
	migrate -database $(POSTGRESQL_URL) -path db/migration up

migrate-down:
	migrate -database $(POSTGRESQL_URL) -path db/migration down