all: build

init:
	@echo "Initializing..."
	@$(MAKE) tool_download
	@$(MAKE) build

build:
	@echo "Building..."
	@go mod tidy
	@$(MAKE) tool_update
	@$(MAKE) proto_gen
	@$(MAKE) sqlc_gen
	@$(MAKE) build_alone

build_alone:
	@echo "Building alone..."
	@$(MAKE) linter-golangci
	@go build -tags migrate -o bin/$(shell basename $(PWD)) ./cmd

pushall:
	@docker build -t ghcr.io/escape-ship/paymentsrv:latest .
	@docker push ghcr.io/escape-ship/paymentsrv:latest

proto_gen:
	@echo "Generating proto..."
	@cd proto && \
	buf dep update && \
	buf lint && \
	buf generate

sqlc_gen:
	@echo "Generating sqlc..."
	@sqlc generate -f internal/infras/sqlc/sqlc.yaml

tool_update:
	@echo "Updating tools..."
	@go get -modfile=tools.mod -tool github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go get -modfile=tools.mod -tool github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go get -modfile=tools.mod -tool google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go get -modfile=tools.mod -tool google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go get -modfile=tools.mod -tool github.com/sqlc-dev/sqlc/cmd/sqlc@latest

tool_download:
	@echo "Downloading tools..."
	@go install -modfile=tools.mod tool
	@go install github.com/bufbuild/buf/cmd/buf@latest

run:
	@echo "Running..."
	@./bin/$(shell basename $(PWD))

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

clean:
	rm -f bin/$(shell basename $(PWD))