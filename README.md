# Thread API

A Go-based thread API application for study purposes.

## Prerequisites

- Go 1.19 or higher
- Make (for using Makefile commands)

## Getting Started

### Running the Application

#### Development Mode
Run the application directly:
```bash
make run
```

#### Production Build
Build the binary and run it:
```bash
make build
./bin/thread-api
```

## Development

### Available Make Commands

- `make run` - Run the application directly
- `make build` - Build the application binary
- `make clean` - Clean build artifacts
- `make fmt` - Format Go code
- `make help` - Show available commands

### Code Formatting

Format your code before committing:
```bash
make fmt
```

## Project Structure

```
thread-api/
├── bin/              # Built binaries (created after make build)
├── Makefile          # Build automation
├── README.md         # This file
├── go.mod            # Go module file
└── main.go           # Main application file
```

## Usage

After running the application with `make run`, the API will be available at the configured port.

