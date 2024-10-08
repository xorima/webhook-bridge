.PHONY: test build docs run

fmt:
	go fmt

test: docs
	go test
build: docs test
	go build

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

run: docs test
	GITHUB_HMAC_ENABLED=false go run .

lint:
	docker run --rm -v $(shell pwd):/app -v ~/.cache/golangci-lint/v1.61.0:/root/.cache -w /app golangci/golangci-lint:v1.61.0 golangci-lint run
