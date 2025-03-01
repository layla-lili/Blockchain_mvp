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

## Transaction Commands

The CLI supports various transaction operations with Anvil local blockchain.

### Sending Transactions

1. **Test Transaction** (between first two Anvil accounts):
```bash
go run cmd/blockchain-cli/main.go tx send --test
```
Example output:
```
Test transaction sent successfully!
From: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
To: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
Value: 1 ETH
Hash: 0x...
```

2. **Regular Transaction** (specify recipient and amount):
```bash
go run cmd/blockchain-cli/main.go tx send \
  --to 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
  --value 1000000000000000000
```
Example output:
```
Transaction sent successfully! Hash: 0x...
```

### Transaction Options

- `--to`: Recipient address (required for non-test transactions)
- `--value`: Amount in wei (1 ETH = 1000000000000000000 wei)
- `--data`: Optional transaction data
- `--test`: Send 1 ETH between first two test accounts

### Development Setup

1. Start Anvil in one terminal:
```bash
anvil
```

2. Run transactions in another terminal:
```bash
# Test mode
go run cmd/blockchain-cli/main.go tx send --test

# Regular transaction
go run cmd/blockchain-cli/main.go tx send --to <ADDRESS> --value <WEI_AMOUNT>
```

### Available Test Accounts

Anvil provides 10 test accounts with 10000 ETH each:
```
(0) 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
(1) 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
(2) 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
(3) 0x90F79bf6EB2c4f870365E785982E1f101E93b906
...
```