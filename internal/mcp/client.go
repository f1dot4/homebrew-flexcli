package mcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a minimal MCP client for calling tools via HTTP.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new MCP client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// mcpRequest represents an MCP tool call request.
type mcpRequest struct {
	Method string `json:"method"`
	Params struct {
		Name      string                 `json:"name"`
		Arguments map[string]interface{} `json:"arguments,omitempty"`
	} `json:"params"`
}

// mcpResponse represents an MCP tool call response.
type mcpResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	IsError bool `json:"isError"`
}

// CallTool calls an MCP tool and returns the result.
func (c *Client) CallTool(name string, args map[string]interface{}) (json.RawMessage, error) {
	// Build request
	req := mcpRequest{
		Method: "tools/call",
	}
	req.Params.Name = name
	req.Params.Arguments = args

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", c.baseURL+"/mcp", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-Key", c.apiKey)

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("MCP error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	// Parse MCP response
	var mcpResp mcpResponse
	if err := json.Unmarshal(body, &mcpResp); err != nil {
		return nil, fmt.Errorf("failed to parse MCP response: %w", err)
	}

	// Check for MCP-level error
	if mcpResp.IsError {
		if len(mcpResp.Content) > 0 {
			return nil, fmt.Errorf("MCP tool error: %s", mcpResp.Content[0].Text)
		}
		return nil, fmt.Errorf("MCP tool error (no details)")
	}

	// Extract result from content
	if len(mcpResp.Content) == 0 {
		return nil, fmt.Errorf("empty MCP response")
	}

	// Parse the text content as JSON
	result := json.RawMessage(mcpResp.Content[0].Text)
	return result, nil
}
