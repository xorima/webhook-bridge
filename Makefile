.PHONY: test build docs run

test: docs
	go test ./...

build: docs test
	go build

docs:
	swag init

run: docs test
	go run .
