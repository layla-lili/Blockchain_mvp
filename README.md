# Blockchain_mvp
Building Blockchain MVP Focus on the absolute minimum features required to demonstrate a functional blockchain 

# Blockchain MVP
This repository contains the core blockchain implementation in Rust (`blockchain_core`) and developer tools in Go (`blockchain_tools`).

## Project Structure

```
.
├── blockchain_core/     # Rust implementation of the blockchain node
└── blockchain_tools/    # Go implementation of CLI tools and APIs
    ├── cmd/
    │   ├── blockchain-api/
    │   └── blockchain-cli/
    ├── internal/
    │   ├── cli/
    │   │   ├── commands/
    │   │   ├── formatter/
    │   │   └── config/
    │   └── common/
    └── pkg/
        ├── client/
        ├── types/
        └── utils/
```

## Setup

### Prerequisites
- Go 1.21 or later
- Rust (latest stable)
- Make

### Blockchain Tools (Go)

1. Navigate to the tools directory:
```bash
cd blockchain_tools
```

2. Build the CLI:
```bash
make build
```

3. Available make commands:
```bash
make help        # Show available commands
make build       # Build the blockchain-cli binary
make clean       # Remove build artifacts
make test        # Run tests
make run         # Build and run the blockchain-cli
make dist        # Create distribution packages
```

4. CLI Usage:
```bash
# Show available commands
./bin/blockchain-cli --help

# Get account balance
./bin/blockchain-cli account balance <address>

# List accounts
./bin/blockchain-cli account list

# Check node status
./bin/blockchain-cli node status
```

### Global Flags
```bash
--config string    # Config file (default: $HOME/.blockchain-cli.yaml)
--debug           # Enable debug logging
--format string   # Output format (table, json, yaml)
--rpc-url string  # Blockchain RPC endpoint URL
```

### Configuration
The CLI can be configured using:
- Command-line flags
- Environment variables
- Configuration file (`$HOME/.blockchain-cli.yaml`)

### Blockchain Core (Rust)

1. Navigate to the core directory:
```bash
cd blockchain_core
```

2. Build the project:
```bash
cargo build
```

3. Run tests:
```bash
cargo test
```

## Development

### Directory Structure
- `cmd/`: Main applications
- `internal/`: Internal packages
- `pkg/`: Public packages
- `examples/`: Example applications

### Testing
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests with race detection
make test-race
```

## License

MIT License - see LICENSE file for details
