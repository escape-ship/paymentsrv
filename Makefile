all: build

init:
	@echo "Initializing..."
	@$(MAKE) tool_download
	@$(MAKE) build

build:
	@echo "Building..."
	@go mod tidy
	@$(MAKE) tool_update
	@$(MAKE) sqlc_gen
	@$(MAKE) build_alone

build_alone:
	@echo "Building alone..."
	@$(MAKE) linter-golangci
	@go build -tags migrate -o bin/$(shell basename $(PWD)) ./cmd

pushall:
	@docker build -t ghcr.io/escape-ship/paymentsrv:latest .
	@docker push ghcr.io/escape-ship/paymentsrv:latest

sqlc_gen:
	@echo "Generating sqlc..."
	@sqlc generate -f internal/infras/sqlc/sqlc.yaml

tool_update:
	@echo "Updating tools..."
	@go get -modfile=tools.mod -tool github.com/sqlc-dev/sqlc/cmd/sqlc@latest

tool_download:
	@echo "Downloading tools..."
	@go install github.com/bufbuild/buf/cmd/buf@latest

run:
	@echo "Running..."
	@./bin/$(shell basename $(PWD))

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

clean:
	rm -f bin/$(shell basename $(PWD))