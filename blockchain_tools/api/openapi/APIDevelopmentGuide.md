# Blockchain API Development Guide

## Overview
This guide documents the API development process for the Blockchain MVP project, which combines a Rust blockchain core with Go-based tools and APIs.

## Architecture

### Components
```
blockchain_tools/
├── api/
│   ├── openapi/
│   │   └── blockchain.yaml    # OpenAPI specification
│   └── swagger/
│       └── ui/               # Embedded Swagger UI files
├── cmd/
│   └── blockchain-api/       # API server implementation
├── internal/
│   └── api/
│       ├── handlers/         # API request handlers
│       ├── middleware/       # Custom middleware
│       └── swagger/          # Swagger UI implementation
└── pkg/
    └── client/              # RPC client implementation
```

## Development Workflow

### 1. Update OpenAPI Specification
```yaml
# filepath: api/openapi/blockchain.yaml
openapi: 3.0.0
info:
  title: Blockchain API
  version: 1.0.0
paths:
  /blocks/{number}:
    get:
      summary: Get block by number
      parameters:
        - name: number
          in: path
          required: true
          schema:
            type: integer
```

### 2. Generate API Code
```bash
# Generate server code
oapi-codegen -package api \
  -generate types,server,spec \
  api/openapi/blockchain.yaml \
  > pkg/api/blockchain.gen.go
```

### 3. Implement API Server
```go
// Example handler implementation
func (s *Server) GetBlock(c *gin.Context) {
    number := c.Param("number")
    // Implementation
}
```

## API Documentation

### Swagger UI Setup

1. **Configuration**
```go
// In cmd/blockchain-api/main.go
swaggerCfg := swagger.Config{
    Title:       "Blockchain API",
    SpecURL:     "/openapi.json",
    BasePath:    "/docs",
    Description: "API for interacting with the blockchain",
}
router.GET("/docs/*any", swagger.Handler(swaggerCfg))
```

2. **Access Points**
- Swagger UI: http://localhost:8080/docs/
- OpenAPI Spec: http://localhost:8080/openapi.json

### OpenAPI Specification

```yaml
# filepath: api/openapi/blockchain.yaml
openapi: 3.0.0
info:
  title: Blockchain API
  version: 1.0.0
  description: API for interacting with the blockchain

servers:
  - url: /api/v1
    description: Development server

paths:
  /transactions:
    post:
      summary: Send a new transaction
      requestBody:
        required: true
        content:
          application/json:
            schema:
              oneOf:
                - $ref: '#/components/schemas/TestTransaction'
                - $ref: '#/components/schemas/CustomTransaction'
            examples:
              test:
                summary: Test Transaction
                value:
                  test: true
              custom:
                summary: Custom Transaction
                value:
                  to: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
                  value: "1000000000000000000"

components:
  schemas:
    TestTransaction:
      type: object
      properties:
        test:
          type: boolean
          description: Set to true for test transaction

    CustomTransaction:
      type: object
      required:
        - to
        - value
      properties:
        to:
          type: string
          description: Recipient address
        value:
          type: string
          description: Amount in wei
```

### Testing Swagger UI

1. **Start Local Environment**
```bash
# Terminal 1: Start Anvil
anvil

# Terminal 2: Start API Server
go run cmd/blockchain-api/main.go
```

2. **Using Swagger UI**
- Open http://localhost:8080/docs/
- Click "Try it out" on any endpoint
- Fill in parameters
- Click "Execute"

3. **Example Requests**
```bash
# Test transaction using Swagger
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Content-Type: application/json" \
  -d '{"test": true}'

# Custom transaction using Swagger
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Content-Type: application/json" \
  -d '{
    "to": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
    "value": "1000000000000000000"
  }'
```

### Swagger UI Customization

1. **Custom Template**
```html
// filepath: internal/api/swagger/swagger-ui/index.html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.11.2/swagger-ui.css">
    <style>
        /* Custom styles */
        .swagger-ui .topbar { display: none }
        .example-section { padding: 20px; background: #f8f9fa; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <div class="example-section">
        <h2>Quick Examples</h2>
        {{range .Examples}}
        <div>
            <h3>{{.Title}}</h3>
            <pre>{{.Curl}}</pre>
        </div>
        {{end}}
    </div>
    <script src="https://unpkg.com/swagger-ui-dist@5.11.2/swagger-ui-bundle.js"></script>
    <script>
        window.onload = () => {
            window.ui = SwaggerUIBundle({
                url: "{{.SpecURL}}",
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIBundle.presets.modals
                ],
            });
        };
    </script>
</body>
</html>
```

### Available Endpoints
```
GET    /api/v1/blocks/:number          # Get block by number
GET    /api/v1/blocks/latest          # Get latest block
POST   /api/v1/transactions           # Send transaction
GET    /api/v1/transactions/:hash     # Get transaction
GET    /api/v1/accounts/:address      # Get account
GET    /api/v1/accounts/:address/balance  # Get account balance
GET    /api/v1/node/status           # Get node status
GET    /api/v1/node/peers            # Get peer list
```

## Configuration

### Environment Variables
```bash
BLOCKCHAIN_RPC_URL  # RPC endpoint URL (default: http://localhost:8545)
GIN_MODE            # Gin framework mode (debug/release)
```

### Server Configuration
```go
type Config struct {
    Port     int    `env:"PORT" default:"8080"`
    RPCUrl   string `env:"BLOCKCHAIN_RPC_URL"`
    LogLevel string `env:"LOG_LEVEL" default:"info"`
}
```

## Development Setup

### Prerequisites
```bash
# Install required tools
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

# Setup development environment
make setup
```

### Build and Run
```bash
# Generate API code
make generate-api

# Run API server
make run-api
```

## Testing

### Running Tests
```bash
# Run all tests
make test

# Run specific tests
go test ./internal/api/...
```

### Testing Endpoints
```bash
# Get latest block
curl http://localhost:8080/api/v1/blocks/latest

# Get account balance
curl http://localhost:8080/api/v1/accounts/0x123.../balance
```

## Best Practices

### 1. API Design
- Use RESTful principles
- Version APIs in URL
- Use proper HTTP methods
- Include validation
- Handle errors consistently

### 2. Documentation
- Keep OpenAPI spec updated
- Document all endpoints
- Include example requests/responses
- Document error conditions

### 3. Error Handling
```yaml
# Standard error response
Error:
  type: object
  properties:
    code:
      type: string
    message:
      type: string
    details:
      type: object
```

### 4. Monitoring
- Log all requests
- Track response times
- Monitor error rates
- Set up alerts

## Troubleshooting

### Common Issues
1. Swagger UI 404
   - Ensure swagger-ui files are embedded
   - Check route configuration

2. RPC Connection
   - Verify RPC endpoint
   - Check network connectivity

3. Generated Code
   - Run `make generate-api`
   - Verify OpenAPI spec syntax

## Local Blockchain Options

### 1. Hardhat (Node.js based, but more developer-friendly)
```bash
# Install
npm install --global hardhat

# Create and start local node
npx hardhat node
```

### 2. Anvil (Foundry's local node - Rust based)
```bash
# Install Foundry (Mac)
curl -L https://foundry.paradigm.xyz | bash
foundryup

# Start local node
anvil
```

### 3. Geth Dev Mode (Official Ethereum client)
```bash
# Install via Homebrew
brew install ethereum

# Start in dev mode
geth --dev --http \
  --http.api eth,web3,personal,net \
  --http.corsdomain "*" \
  --http.addr 0.0.0.0 \
  --http.port 8545
```

### 4. LocalEVM (Lightweight Go implementation)
```bash
# Install
go install github.com/ethereumjs/ethereumjs-monorepo/packages/local-evm@latest

# Start
localevm
```

### Local Node Configuration
```yaml
# Common Settings for All Options
RPC URL: http://localhost:8545
Chain ID: 1337
Network ID: 1337
Default Gas Limit: 6721975
Default Gas Price: 20000000000
```

### Test Accounts (Common across implementations)
```
Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Balance: 10000 ETH

Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
Balance: 10000 ETH
```

### Switching to Rust Blockchain Core

Once your Rust blockchain core is ready:

1. Update RPC URL to point to your Rust node
```bash
export BLOCKCHAIN_RPC_URL=http://localhost:9545  # Your Rust node port
```

2. Update RPC client implementation if needed
```go
// filepath: pkg/client/rpc/client.go
// Implement any custom RPC methods specific to your blockchain

## References
- [OpenAPI Specification](https://swagger.io/specification/)
- [Gin Web Framework](https://gin-gonic.com/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)