package api

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/getkin/kin-openapi/openapi3"
)

// GetSwagger returns the Swagger specification loaded from blockchain.yaml
func GetSwagger() *openapi3.T {
	// Get the directory of the current file
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "../..")

	// Construct the path to the OpenAPI spec
	specPath := filepath.Join(rootDir, "api", "openapi", "blockchain.yaml")

	// Load and parse the spec
	swagger, err := openapi3.NewLoader().LoadFromFile(specPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load swagger spec from %s: %v", specPath, err))
	}

	return swagger
}
