package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLevel represents the severity of the log entry
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

// generateRequestID creates a unique request ID
func generateRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "unknown"
	}
	return hex.EncodeToString(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID
		requestID := generateRequestID()
		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)

		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Determine log level based on status code
		var level LogLevel
		status := c.Writer.Status()
		switch {
		case status >= 500:
			level = ERROR
		case status >= 400:
			level = WARN
		default:
			level = INFO
		}

		// Collect errors if any
		errors := c.Errors.String()

		// Create structured log entry
		logEntry := struct {
			Timestamp string `json:"timestamp"`
			Level     string `json:"level"`
			RequestID string `json:"request_id"`
			Method    string `json:"method"`
			Path      string `json:"path"`
			Status    int    `json:"status"`
			Duration  string `json:"duration"`
			ClientIP  string `json:"client_ip"`
			Errors    string `json:"errors,omitempty"`
		}{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Level:     level.String(),
			RequestID: requestID,
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    status,
			Duration:  time.Since(start).String(),
			ClientIP:  c.ClientIP(),
			Errors:    errors,
		}

		// Build log message using strings.Builder
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("[%s] ", level))
		sb.WriteString(fmt.Sprintf("RequestID=%s ", requestID))
		sb.WriteString(fmt.Sprintf("Method=%s ", c.Request.Method))
		sb.WriteString(fmt.Sprintf("Path=%s ", c.Request.URL.Path))
		sb.WriteString(fmt.Sprintf("Status=%d ", status))
		sb.WriteString(fmt.Sprintf("Duration=%s ", time.Since(start)))
		sb.WriteString(fmt.Sprintf("ClientIP=%s", c.ClientIP()))

		if errors != "" {
			sb.WriteString(fmt.Sprintf(" Errors=%s", errors))
		}
		sb.WriteString("\n")

		// Write log entry with error handling
		if _, err := gin.DefaultWriter.Write([]byte(sb.String())); err != nil {
			log.Printf("Failed to write log entry: %v", err)
		}

		// Store structured log entry in context for potential later use
		c.Set("LogEntry", logEntry)
	}
}
