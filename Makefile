APP_NAME=lobster
MAIN_PATH=./cmd/api/main.go

.PHONY: run dev build

run:
	go run $(MAIN_PATH)

dev:
	air

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)
