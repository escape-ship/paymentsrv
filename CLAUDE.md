# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Build and Development
- `make build` - Full build with proto generation, sqlc generation, and binary compilation
- `make build_alone` - Build binary only (includes linting)
- `make run` - Run the compiled binary
- `make linter-golangci` - Run golangci-lint for code quality checks
- `make test` - Run tests (as documented in README)

### Code Generation
- `make proto_gen` - Generate protobuf files using buf
- `make sqlc_gen` - Generate SQL code using sqlc
- `make tool_download` - Download required tools
- `make tool_update` - Update tools to latest versions

### Development Workflow
1. Run `make init` for first-time setup
2. Use `make build` for full rebuilds
3. Use `make build_alone` for quick builds after changes
4. Always run `make linter-golangci` before committing

## Architecture Overview

This is a Go-based payment microservice with gRPC API, integrating with Kakao payments and using Kafka for event streaming.

### Key Components

**Application Layer** (`internal/app/`):
- `app.go` - Main application struct with gRPC server setup, graceful shutdown, and dependency injection

**Service Layer** (`internal/service/`):
- `service.go` - gRPC service implementation for payment operations (KakaoReady, KakaoApprove, KakaoCancel)
- Handles business logic, database transactions, and Kafka message publishing

**Infrastructure Layer**:
- `internal/infras/sqlc/` - Database layer using sqlc for type-safe SQL queries
- `internal/provider/kakao/` - Kakao payment provider integration
- `pkg/kafka/` - Kafka producer/consumer interfaces and implementation
- `pkg/postgres/` - PostgreSQL database connection and interface

**Configuration** (`config/`):
- Uses Viper for config management with YAML files and environment variables
- Supports app, database, and Kakao payment settings

### Database Management
- Uses golang-migrate for database migrations (see `db/migrations/`)
- SQLC generates type-safe Go code from SQL queries
- PostgreSQL with transaction support

### Message Streaming
- Kafka integration for event-driven architecture
- Publishes messages on successful payment approvals
- Configurable producer/consumer interfaces

### External Integrations
- Kakao Pay API for payment processing
- Protocol Buffers for gRPC service definitions
- Buf for protobuf management and code generation

## Project Structure Notes

- Binary output: `bin/` directory
- Proto definitions: `proto/` with generated code in `proto/gen/`
- Database queries: `internal/infras/sqlc/query.sql`
- Main entry point: `cmd/main.go`
- Configuration: `config.yaml` with environment variable overrides

## Testing

The project uses `make test` for running tests. Test files should follow standard Go testing conventions.