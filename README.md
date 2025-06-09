# Go Web API Sample

A modern Go web API sample project using TypeSpec for API specification and ogen for code generation. This project demonstrates a clean, layered architecture with comprehensive CRUD operations for note management.

## Features

- **Type-safe API**: TypeSpec specifications with automatic Go code generation
- **Clean Architecture**: 3-layer design (handler, service, repository)
- **CRUD Operations**: Complete note management (create, read, update, delete)
- **OpenAPI Integration**: Automatic OpenAPI 3.1 schema generation
- **Memory Storage**: In-memory repository implementation (easily extensible)

## API Endpoints

- `GET /notes` - List all notes
- `GET /notes/{id}` - Get a specific note by ID
- `POST /notes` - Create a new note
- `DELETE /notes/{id}` - Delete a note by ID

## Quick Start

### Prerequisites

- Go 1.24.2 or later
- Node.js (for TypeSpec compilation)
- Task (task runner)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   task init
   ```

### Development Workflow

1. **Edit API specifications** in `spec/src/`
2. **Generate code** from specifications:
   ```bash
   task spec:generate
   ```
3. **Run the server**:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

### Example Usage

```bash
# Create a note
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "title": "Sample Note", "content": "This is a sample note"}'

# List all notes
curl http://localhost:8080/notes

# Get a specific note
curl http://localhost:8080/notes/1

# Delete a note
curl -X DELETE http://localhost:8080/notes/1
```

## Project Structure

```
├── internal/          # Application code (not importable by external packages)
│   ├── handler/       # HTTP handlers
│   ├── service/       # Business logic
│   └── repository/    # Data access layer
│       └── memory/    # In-memory implementation
├── ogen/             # Generated Go code (auto-generated)
├── spec/             # TypeSpec API specifications
│   └── src/
│       ├── models/    # Data models
│       └── routes/    # API routes
├── schema/           # Generated OpenAPI schema (auto-generated)
├── main.go           # Application entry point
└── CLAUDE.md         # Development guidance
```

## Architecture

This project follows a **specification-first approach**:

1. **TypeSpec Specifications** (`spec/src/`) - Define API contracts using TypeScript-like syntax
2. **OpenAPI Generation** - Compile TypeSpec to OpenAPI 3.1 schema
3. **Go Code Generation** - Generate type-safe Go code using ogen
4. **Implementation** - Implement business logic using generated interfaces

### Key Technologies

- **[TypeSpec](https://typespec.io/)**: API specification language
- **[ogen](https://github.com/ogen-go/ogen)**: OpenAPI Go code generator
- **[Task](https://taskfile.dev/)**: Task runner for build automation

## Development Commands

```bash
# Initialize dependencies
task init

# Compile TypeSpec and generate Go code
task spec:generate

# Build the application
go build

# Run the application
go run main.go

# Run with Docker
docker build -t go-web-api-sample .
docker run -p 8080:8080 go-web-api-sample
```

## Extending the Project

### Adding New APIs

1. Edit TypeSpec files in `spec/src/`
2. Run `task spec:generate` to regenerate code
3. Implement the new interfaces in your application code

### Adding New Storage Backends

Implement the `repository.NoteRepository` interface:

```go
type NoteRepository interface {
    Create(ctx context.Context, note *ogen.ModelsNote) (*ogen.ModelsNote, error)
    GetByID(ctx context.Context, id int64) (*ogen.ModelsNote, error)
    List(ctx context.Context) ([]ogen.ModelsNote, error)
    Delete(ctx context.Context, id int64) error
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `task spec:generate` if you modified TypeSpec files
5. Test your changes
6. Submit a pull request

## License

This project is provided as a sample for educational purposes.