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

### Swagger UI Access
- Local Development: http://localhost:8080/docs/
- OpenAPI Spec: http://localhost:8080/openapi.json

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

## References
- [OpenAPI Specification](https://swagger.io/specification/)
- [Gin Web Framework](https://gin-gonic.com/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)