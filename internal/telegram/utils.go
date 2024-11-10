package telegram

import (
	"fmt"
	"strconv"
	"strings"
)

// parseIntParams parses the input parameters for the command.
func parseIntParams(text string, requiredKeys []string) (map[string]int, error) {
	params := map[string]int{}
	pairs := strings.Fields(text)[1:] // Skip the command part

	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid parameter format")
		}
		key := kv[0]
		value, err := strconv.Atoi(kv[1])
		if err != nil {
			//return nil, fmt.Errorf("invalid integer value for %s", key)
			continue
		}
		params[key] = value
	}

	// Check that all required parameters are present
	for _, key := range requiredKeys {
		if _, exists := params[key]; !exists {
			return nil, fmt.Errorf("missing required parameter: %s", key)
		}
	}

	return params, nil
}

// parseStrParams parses the input parameters for the command.
func parseStrParams(text string, requiredKeys []string) (map[string]string, error) {
	params := map[string]string{}
	pairs := strings.Fields(text)[1:] // Skip the command part

	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid parameter format")
		}
		key := kv[0]
		value := kv[1]

		params[key] = value
	}

	// Check that all required parameters are present
	for _, key := range requiredKeys {
		if _, exists := params[key]; !exists {
			return nil, fmt.Errorf("missing required parameter: %s", key)
		}
	}

	return params, nil
}
