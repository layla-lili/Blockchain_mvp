package swagger

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// Embed the swagger-ui directory
//
//go:embed swagger-ui/*
var swaggerUI embed.FS

// Config holds Swagger UI configuration
type Config struct {
	Title       string
	SpecURL     string
	BasePath    string
	Description string
	Examples    []Example
}

// Example represents a transaction example
type Example struct {
	Title string
	Curl  string
}

// Handler returns a gin.HandlerFunc that serves the Swagger UI
func Handler(config Config) gin.HandlerFunc {
	// Add transaction examples to config
	config.Examples = []Example{
		{
			Title: "Send Test Transaction",
			Curl: `curl -X POST http://localhost:8080/api/v1/transactions \
				-H "Content-Type: application/json" \
				-d '{"test": true}'`,
		},
		{
			Title: "Send Custom Transaction",
			Curl: `curl -X POST http://localhost:8080/api/v1/transactions \
				-H "Content-Type: application/json" \
				-d '{
					"to": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
					"value": "1000000000000000000"
				}'`,
		},
	}

	return func(c *gin.Context) {
		log.Printf("Swagger request path: %s", c.Param("any"))

		// If requesting the root docs path, serve the HTML
		if c.Param("any") == "/" || c.Param("any") == "" {
			tmpl, err := template.ParseFS(swaggerUI, "swagger-ui/index.html")
			if err != nil {
				log.Printf("Failed to parse template: %v", err)
				c.String(http.StatusInternalServerError, "Failed to load Swagger UI template")
				return
			}

			c.Header("Content-Type", "text/html")
			err = tmpl.Execute(c.Writer, config)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to render Swagger UI template")
			}
			return
		}

		// For other paths, try to serve static files
		filepath := path.Join("swagger-ui", c.Param("any"))
		c.FileFromFS(filepath, http.FS(swaggerUI))
	}
}

// GetSwagger returns the parsed OpenAPI specification
// Don't hardcode the spec in Go files
// Instead, load from blockchain.yaml:
func GetSwagger() *openapi3.T {
	specPath := "api/openapi/blockchain.yaml"
	swagger, err := openapi3.NewLoader().LoadFromFile(specPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load spec: %v", err))
	}
	return swagger
}

func LoadSwagger(path string) (*openapi3.T, error) {
	swagger, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load swagger spec: %w", err)
	}
	return swagger, nil
}
