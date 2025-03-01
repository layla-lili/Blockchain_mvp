# Makefile for blockchain tools

# Variables
BINARY_NAME=blockchain-cli
GO=go
VERSION := $(shell git describe --tags --always --dirty)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date -u '+%Y-%m-%d-%H:%M-%Z')

# Build flags
LDFLAGS=-ldflags "-X github.com/layla-lili/blockchain_tools/internal/cli/commands.Version=$(VERSION) -X github.com/layla-lili/blockchain_tools/internal/cli/commands.GitCommit=$(COMMIT) -X github.com/layla-lili/blockchain_tools/internal/cli/commands.BuildDate=$(DATE)"

.PHONY: all build clean run test lint

# Default target
all: build

# Build the application
build:
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p bin
	go build -o bin/$(BINARY_NAME) $(LDFLAGS) ./cmd/blockchain-cli

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf dist/

# Run the application
run: build
	@echo "Running ${BINARY_NAME}..."
	@chmod +x bin/${BINARY_NAME}
	./bin/${BINARY_NAME}

# Run tests
test:
	@echo "Running tests..."
	$(GO) test ./... -v

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Create distribution packages
dist:
	@echo "Creating distribution packages..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/${BINARY_NAME}-linux-amd64 ./cmd/blockchain-cli
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/${BINARY_NAME}-darwin-amd64 ./cmd/blockchain-cli
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) $(LDFLAGS) -o dist/${BINARY_NAME}-windows-amd64.exe ./cmd/blockchain-cli

# Help
help:
	@echo "Blockchain Tools Makefile"
	@echo "Available targets:"
	@echo "  build     - Build the blockchain-cli binary"
	@echo "  install   - Install the blockchain-cli binary"
	@echo "  clean     - Remove build artifacts"
	@echo "  test      - Run tests"
	@echo "  lint      - Run linter"
	@echo "  run       - Build and run the blockchain-cli"
	@echo "  dist      - Create distribution packages"
