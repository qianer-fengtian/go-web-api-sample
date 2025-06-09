# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Architecture

This is a Go web API sample project using **TypeSpec for API specification** and **ogen for code generation**. The architecture follows a dual-specification approach:

### Core Components
- **TypeSpec specifications** in `spec/src/` define the API contract using TypeScript-like syntax
- **Generated OpenAPI** schema in `schema/openapi.yaml` (auto-generated from TypeSpec)
- **Generated Go code** in `ogen/` directory (auto-generated from OpenAPI schema)
- **Application logic** in `main.go` implements the business logic

### Development Workflow
The project uses a specification-first approach:
1. Edit TypeSpec files in `spec/src/`
2. Compile TypeSpec to OpenAPI YAML
3. Generate Go code from OpenAPI schema
4. Implement business logic using generated interfaces

## Common Commands

### Build and Code Generation
```bash
# Initialize dependencies (required first time)
task init

# Compile TypeSpec specifications to OpenAPI YAML
task spec:compile

# Generate Go code from OpenAPI schema  
task spec:generate

# Full regeneration workflow
task spec:generate
```

### Development
```bash
# Run the application
go run main.go

# Build the application
go build

# Install Go dependencies
go mod tidy
```

## Code Generation Details

- **TypeSpec source**: `spec/src/main.tsp` and subdirectories
- **Generated OpenAPI**: `schema/openapi.yaml` 
- **Generated Go interfaces**: `ogen/oas_handlers_gen.go`
- **ogen configuration**: `ogen.yml` (ignores unsupported content types)

The generated Go code provides type-safe interfaces that must be implemented in your application code. The `main.go` file shows how to implement these interfaces and wire up the HTTP server.

## Architecture Notes

- TypeSpec uses namespaces (`NoteService.Models`, `NoteService.Routes`) for organization
- Generated code is completely replaced on each run - never edit files in `ogen/`
- Application logic should implement the generated interfaces rather than defining custom HTTP handlers
- The project uses Go modules with toolchain dependencies for ogen and task