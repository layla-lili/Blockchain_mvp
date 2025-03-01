package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/layla-lili/blockchain_tools/internal/api"
	"github.com/layla-lili/blockchain_tools/internal/api/handlers"
	"github.com/layla-lili/blockchain_tools/internal/api/middleware"
	"github.com/layla-lili/blockchain_tools/internal/api/swagger"
	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	// Initialize RPC client
	rpcURL := os.Getenv("BLOCKCHAIN_RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://localhost:8545" // default value
	}
	client, err := rpc.NewClient(rpcURL)
	if err != nil {
		log.Fatalf("Failed to create RPC client: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Load API spec
	spec := api.GetSwagger()
	if spec == nil {
		log.Fatal("Failed to load OpenAPI specification")
	}

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Serve Swagger UI
	swaggerCfg := swagger.Config{
		Title:    "Blockchain API",
		SpecURL:  "/openapi.json",
		BasePath: "/docs",
	}
	router.GET("/docs/*any", swagger.Handler(swaggerCfg))

	// Serve OpenAPI spec
	router.GET("/openapi.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, spec)
	})

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Cors())

	router.Use(func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": c.Errors.Errors(),
			})
		}
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Block endpoints
		api.GET("/blocks/:number", handlers.GetBlock(client))
		api.GET("/blocks/latest", handlers.GetLatestBlock(client))

		// Transaction endpoints
		api.POST("/transactions", handlers.SendTransaction(client))
		api.GET("/transactions/:hash", handlers.GetTransaction(client))

		// Account endpoints
		api.GET("/accounts/:address", handlers.GetAccount(client))
		api.GET("/accounts/:address/balance", handlers.GetBalance(client))

		// Node info endpoints
		api.GET("/node/status", handlers.GetNodeStatus(client))
		api.GET("/node/peers", handlers.GetPeers(client))
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}()

	// Start server
	log.Printf("Starting API server on :8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}
