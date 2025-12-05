.PHONY: build run clean fmt help

# Variables
BINARY_NAME=thread-api

# Build the application
build:
	go build -o bin/$(BINARY_NAME) .

# Run the application directly
run:
	go run .

# Clean build artifacts
clean:
	go clean
	rm -rf bin/

# Format code
fmt:
	go fmt ./...

# Help
help:
	@echo "Available targets:"
	@echo "  build    Build the application"
	@echo "  run      Run the application"
	@echo "  clean    Clean build artifacts"
	@echo "  fmt      Format code"
	@echo "  help     Show this help"