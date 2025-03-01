# Blockchain Tools

Go-based CLI and developer tools for interacting with the blockchain.

## Project Structure

This project follows the standard Go project layout:

- `cmd/`: Contains the main applications
- `internal/`: Private application and library code
- `pkg/`: Library code that's ok to use by external applications
- `api/`: API definitions
- `examples/`: Example applications
- `scripts/`: Scripts for development
- `docs/`: Documentation

.
├── cmd/
│   ├── blockchain-api/    # HTTP API server
│   └── blockchain-cli/    # Command line interface
├── internal/
│   ├── cli/              # CLI implementation
│   │   ├── commands/     # CLI commands
│   │   ├── formatter/    # Output formatters
│   │   └── config/       # Configuration
│   └── common/           # Shared utilities
└── pkg/
    ├── client/           # Blockchain client
    ├── types/            # Common types
    └── utils/ 

## Getting Started

### Prerequisites

- Go 1.21 or later
- Make

### Build & Run

```bash
# Build the CLI
make build

# Run the CLI
./bin/blockchain-cli --help
```

### Available Commands

```bash
blockchain-cli [command]

Available Commands:
  account     Manage blockchain accounts
  block       Manage blockchain blocks
  node        Manage blockchain node
  tx          Manage transactions
  version     Show version information
  help        Help about any command
```

### Global Flags

```bash
--config string    # Config file location
--debug           # Enable debug logging
--format string   # Output format (table, json, yaml)
--rpc-url string  # RPC endpoint URL
```

### Make Commands

```bash
make help         # Show available commands
make build        # Build the binary
make clean        # Remove build artifacts
make test         # Run tests
make run          # Build and run
make dist         # Create distribution
```

## Development

### Running Tests

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run with race detection
make test-race
```

### Code Structure

- Commands are in `internal/cli/commands/`
- RPC client is in `pkg/client/rpc/`
- Types are in `pkg/types/`