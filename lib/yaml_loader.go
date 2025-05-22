package lib

import (
	"os"
	"gopkg.in/yaml.v3"
	"fmt"
)

// LoadYAML reads a YAML file from the given filePath and unmarshals it into a map.
func LoadYAML(filePath string) (map[string]interface{}, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML content
	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetString safely extracts a string value from a map[string]interface{}.
// It returns the defaultValue if the key is not found, the value is nil, or the type is not string.
func GetString(data map[string]interface{}, key string, defaultValue string) string {
	if val, ok := data[key]; ok && val != nil {
		if strVal, typeOk := val.(string); typeOk {
			return strVal
		}
		fmt.Printf("Warning: Type mismatch for key '%s'. Expected string, got %T. Using default value '%s'.\n", key, val, defaultValue)
	}
	return defaultValue
}

// GetInt safely extracts an integer value from a map[string]interface{}.
// It returns the defaultValue if the key is not found, the value is nil, or the type is not int.
// It also handles cases where the number might be a float64 (common in YAML parsing) and attempts conversion.
func GetInt(data map[string]interface{}, key string, defaultValue int) int {
	if val, ok := data[key]; ok && val != nil {
		switch v := val.(type) {
		case int:
			return v
		case float64: // YAML parsers often decode numbers as float64
			return int(v)
		case float32:
			return int(v)
		default:
			fmt.Printf("Warning: Type mismatch for key '%s'. Expected int or float, got %T. Using default value %d.\n", key, val, defaultValue)
		}
	}
	return defaultValue
}

// GetBool safely extracts a boolean value from a map[string]interface{}.
// It returns the defaultValue if the key is not found, the value is nil, or the type is not bool.
func GetBool(data map[string]interface{}, key string, defaultValue bool) bool {
	if val, ok := data[key]; ok && val != nil {
		if boolVal, typeOk := val.(bool); typeOk {
			return boolVal
		}
		fmt.Printf("Warning: Type mismatch for key '%s'. Expected bool, got %T. Using default value %t.\n", key, val, defaultValue)
	}
	return defaultValue
}
