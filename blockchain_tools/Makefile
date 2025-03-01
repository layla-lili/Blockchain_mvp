# Makefile for blockchain tools

# Variables
BINARY_NAME=blockchain-cli
GO=go
VERSION := $(shell git describe --tags --always --dirty)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date -u '+%Y-%m-%d-%H:%M-%Z')

# Build flags
LDFLAGS=-ldflags "-X github.com/layla-lili/blockchain_tools/internal/cli/commands.Version=$(VERSION) -X github.com/layla-lili/blockchain_tools/internal/cli/commands.GitCommit=$(COMMIT) -X github.com/layla-lili/blockchain_tools/internal/cli/commands.BuildDate=$(DATE)"

.PHONY: all build clean run test lint download-swagger-ui start-local dev-env tx-test tx-send

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

# Download Swagger UI
download-swagger-ui:
	@echo "Downloading Swagger UI..."
	@mkdir -p internal/api/swagger/swagger-ui
	@curl -L https://github.com/swagger-api/swagger-ui/archive/refs/tags/v5.11.2.tar.gz | \
		tar xz -C /tmp && \
		cp -r /tmp/swagger-ui-5.11.2/dist/* internal/api/swagger/swagger-ui/

# Start local environment
start-local: start-node start-api
	@echo "Testing connection..."
	@sleep 2
	@make tx-test

start-node:
	@echo "Starting Anvil node..."
	anvil --chain-id 1337 --block-time 2

start-api:
	@echo "Starting API server..."
	cd cmd/blockchain-api && go run .

start-cli:
	@echo "Starting CLI..."
	cd cmd/blockchain-cli && go run . --rpc-url http://localhost:8545

# Transaction commands
tx-test:
	@echo "Sending test transaction..."
	./bin/$(BINARY_NAME) tx send --test

tx-send:
	@echo "Usage: make tx-send TO=<address> VALUE=<wei_amount>"
	@if [ -n "$(TO)" ] && [ -n "$(VALUE)" ]; then \
		./bin/$(BINARY_NAME) tx send --to $(TO) --value $(VALUE); \
	fi

# Help
help:
	@echo "Blockchain Tools Makefile"
	@echo "Available targets:"
	@echo "  build       - Build the blockchain-cli binary"
	@echo "  tx-test    - Send a test transaction between first two accounts"
	@echo "  tx-send    - Send a transaction (requires TO and VALUE)"
	@echo "  clean       - Remove build artifacts"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linter"
	@echo "  run         - Build and run the blockchain-cli"
	@echo "  dist        - Create distribution packages"
	@echo "  start-local - Start local environment with Anvil"
