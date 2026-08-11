package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// JsonParser struct
type JsonParser struct{}

// Parses the JSON file and returns a generic map representation of the JSON data
func (j *JsonParser) Parse(filePath string) (map[string]any, error) {
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Unmarshal JSON
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
