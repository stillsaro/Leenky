include .env

# Go command
build:
	@go build -o ./bin/Leenky ./cmd/main/

run: build
	@./bin/Leenky

# Air config command
dev:
	@air --build.cmd "go build -o bin/Leenky cmd/main/" --build.bin "./bin/Leenky"

# Database command
DBInit:
	@docker run --name postgres-container -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d -p 5432:5432 postgres

DBLoad:
	@docker restart postgres-container

DB:
	@docker exec -it postgres-container psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

gooseUp:
	@goose postgres postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB} -dir ./internal/database/schema up

gooseDown:
	@goose postgres postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB} -dir ./internal/database/schema down

sqlc:
	@sqlc generate
