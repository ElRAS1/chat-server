include .env
LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
.PHONY: install-deps

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml
.PHONY: lint

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
.PHONY: install-deps

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
.PHONY: get-deps

generate:
	make generate-chatServer-api
.PHONY: generate

generate-chatServer-api:
	mkdir -p pkg/chatServer
	protoc --proto_path api/chatServer \
	--go_out=pkg/chatServer --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chatServer --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chatServer/chatServer.proto
.PHONY: generate-chatServer-api

format:
	go fmt ./...
.PHONY: format

run:
	go run cmd/main.go
.PHONY: run

migrate-down:
	migrate -path=internal/migrations -database=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable down
.PHONY: migrate-down

migrate-up:
	migrate -path=internal/migrations -database=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable up
.PHONY: migrate-ip
