package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// YamlParser struct
type YamlParser struct{}

// Parses the YAML file and returns a generic map representation of the YAML data
func (y *YamlParser) Parse(filePath string) (map[string]any, error) {
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	// Unmarshal YAML
	var result map[string]any
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return result, nil
}
