package formatters

import (
	"encoding/json"
	"fmt"
)

// RenderJSON outputs data as formatted JSON.
func RenderJSON(data interface{}) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(output))
	return nil
}
