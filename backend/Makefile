.PHONY: build run test fmt lint

BINARY_NAME=personal-website-server

build:
	go build -o bin/${BINARY_NAME} ./cmd/app/main.go

run:
	go run cmd/app/main.go

test:
	go test ./... -v

fmt:
	go fmt ./...

lint:
	golangci-lint run