include .env
export

LOCAL_BIN:=$(CURDIR)/bin

MIGRATION_FOLDER=$(CURDIR)/internal/pkg/database/migrations

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

generate:
	make install-deps
	make get-deps
	make generate-auth-api

generate-auth-api:
	mkdir -p internal/pb/auth
	protoc --proto_path api/auth --proto_path vendor.protogen \
	       --proto_path=vendor.protogen/google \
	       --go_out=internal/pb/auth --go_opt=paths=source_relative \
	       --plugin=protoc-gen-go=bin/protoc-gen-go \
	       --go-grpc_out=internal/pb/auth --go-grpc_opt=paths=source_relative \
	       --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	       --grpc-gateway_out=internal/pb/auth --grpc-gateway_opt=paths=source_relative \
	       --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	       api/auth/auth.proto

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable" up
