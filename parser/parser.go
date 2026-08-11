package parser

import (
	"fmt"
	"strings"
)

// Parser interface
type Parser interface {
	Parse(filePath string) (map[string]any, error)
}

// NewParser returns a new parser instance based on the file extension
func NewParser(filePath string) (Parser, error) {
	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		return &YamlParser{}, nil
	} else if strings.HasSuffix(filePath, ".json") {
		return &JsonParser{}, nil
	}
	return nil, fmt.Errorf("unsupported input format")
}

// Comfort function to parse input file and return a generic map representation of the data
func ParseInput(filePath string) (map[string]any, error) {
	parser, err := NewParser(filePath)
	if err != nil {
		return nil, err
	}

	return parser.Parse(filePath)
}
