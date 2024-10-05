.PHONY: test build docs run

test: docs
	go test ./...

build: docs test
	go build

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

run: docs test
	go run .
