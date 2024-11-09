include .env
LOCAL_BIN:=$(CURDIR)/bin

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml --fix
.PHONY: lint

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
PHONY: install-deps

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2
.PHONY: get-deps

generate:
	make generate-chatServer-api
.PHONY: generate

generate-chatServer-api:
	mkdir -p pkg/chatServer
	mkdir -p pkg/swagger
	protoc --proto_path api/chatServer --proto_path vendor.protogen \
	--go_out=pkg/chatServer --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chatServer --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/chatServer --grpc-gateway_opt=paths=source_relative \
    --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
    --openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
    --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/chatServer/chatServer.proto
.PHONY: generate-chatServer-api

vendor-proto:
		if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi
PHONY: vendor-proto

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
