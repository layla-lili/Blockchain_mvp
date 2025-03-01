package logging

import (
	"io"
	"log"
	"os"
	"sync"
)

// Logger represents a simple logging interface
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
}

// LogLevel represents the logging level
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelError
)

type LogFormat struct {
	TimeFormat string
	Prefix     string
	Flags      int
}

type defaultLogger struct {
	mu          sync.RWMutex
	level       LogLevel
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

var globalLogger = &defaultLogger{
	level:       LevelInfo, // Default level
	debugLogger: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
	infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
	errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
}

func init() {
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		SetLevel(level)
	}
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	return globalLogger
}

// SetLevel sets the global logging level
func SetLevel(level string) {
	globalLogger.mu.Lock()
	defer globalLogger.mu.Unlock()

	switch level {
	case "debug":
		globalLogger.level = LevelDebug
	case "info":
		globalLogger.level = LevelInfo
	case "error":
		globalLogger.level = LevelError
	default:
		globalLogger.level = LevelInfo
	}
}

func SetOutput(w io.Writer) {
	globalLogger.mu.Lock()
	defer globalLogger.mu.Unlock()
	globalLogger.debugLogger.SetOutput(w)
	globalLogger.infoLogger.SetOutput(w)
	globalLogger.errorLogger.SetOutput(w)
}

func (l *defaultLogger) shouldLog(msgLevel LogLevel) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return msgLevel >= l.level
}

func (l *defaultLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.shouldLog(LevelInfo) {
		l.infoLogger.Printf("%s %v\n", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *defaultLogger) Error(msg string, keysAndValues ...interface{}) {
	if l.shouldLog(LevelError) {
		l.errorLogger.Printf("%s %v\n", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *defaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.shouldLog(LevelDebug) {
		l.debugLogger.Printf("%s %v\n", msg, formatKeyValues(keysAndValues...))
	}
}

// formatKeyValues formats key-value pairs into a string
func formatKeyValues(keysAndValues ...interface{}) string {
	if len(keysAndValues) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i]
		var value interface{} = "<no value>"
		if i+1 < len(keysAndValues) {
			value = keysAndValues[i+1]
		}
		result += " " + key.(string) + "=" + value.(string)
	}
	return result
}
