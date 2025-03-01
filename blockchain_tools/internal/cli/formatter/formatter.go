// internal/cli/formatter/formatter.go
package formatter

import (
	"io"
)

// Formatter defines an interface for output formatting
type Formatter interface {
	Format(w io.Writer, data interface{}) error
}

// GetFormatter returns a formatter for the specified format
func GetFormatter(format string) Formatter {
	switch format {
	case "json":
		return NewJSONFormatter()
	case "yaml":
		return NewYAMLFormatter()
	default:
		return NewTableFormatter()
	}
}