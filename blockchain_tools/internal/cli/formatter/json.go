package formatter

import (
	"encoding/json"
	"io"
)

// JSONFormatter implements the Formatter interface for JSON output
type JSONFormatter struct {
	// We can add configuration options here if needed
	// like indentation preference
	indent bool
}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{
		indent: true, // Default to pretty-printed JSON
	}
}

// Format formats the data as JSON and writes it to the writer
func (f *JSONFormatter) Format(w io.Writer, data interface{}) error {
	encoder := json.NewEncoder(w)

	if f.indent {
		encoder.SetIndent("", "  ")
	}

	return encoder.Encode(data)
}
