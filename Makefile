.PHONY: build-app

# Build commands
build-app:
	go build  -o ./bin/app ./cmd/app/main.go

build-migrator-up:
	go build  -o ./bin/migrator/up ./cmd/migrator/up/main.go

# Run commands
run-local:
	CONFIG_PATH=./config/local.yaml ./bin/app

run-prod:
	CONFIG_PATH=./config/prod.yaml ./bin/app

run-migrate-up-local:
	CONFIG_PATH=./config/local.yaml ./bin/migrator/up

run-migrate-up-prod:
	CONFIG_PATH=./config/prod.yaml ./bin/migrator/up


.DEFAULT_GOAL := build-app