package formatter

import (
	"io"

	"gopkg.in/yaml.v2"
)

// YAMLFormatter implements the Formatter interface for YAML output
type YAMLFormatter struct{}

// NewYAMLFormatter creates a new YAML formatter
func NewYAMLFormatter() *YAMLFormatter {
	return &YAMLFormatter{}
}

// Format formats the data as YAML and writes it to the writer
func (f *YAMLFormatter) Format(w io.Writer, data interface{}) error {
	// Marshal the data to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// Write the YAML data to the writer
	_, err = w.Write(yamlData)
	return err
}
