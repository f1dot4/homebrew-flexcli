package mcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Client is a minimal MCP client for calling tools via HTTP.
type Client struct {
	baseURL     string
	apiKey      string
	sessionID   string
	initialized bool
	mu          sync.Mutex
	httpClient  *http.Client
}

// NewClient creates a new MCP client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL:     baseURL,
		apiKey:      apiKey,
		initialized: false,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// mcpRequest represents an MCP JSON-RPC 2.0 request.
type mcpRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// mcpResponse represents an MCP JSON-RPC 2.0 response.
type mcpResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *mcpError       `json:"error,omitempty"`
}

// mcpError represents an MCP JSON-RPC 2.0 error.
type mcpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ensureInitialized performs the MCP initialize handshake if needed.
func (c *Client) ensureInitialized() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.initialized {
		return nil
	}

	// Build initialize request
	req := mcpRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":    map[string]interface{}{},
			"clientInfo": map[string]interface{}{
				"name":    "flexcli",
				"version": "1.0.0",
			},
		},
	}

	// Send initialize request
	resp, sessionID, err := c.sendRequest(req, "")
	if err != nil {
		return fmt.Errorf("initialize failed: %w", err)
	}

	// Check for error response
	if resp.Error != nil {
		return fmt.Errorf("initialize error: %s", resp.Error.Message)
	}

	// Store session ID if provided
	if sessionID != "" {
		c.sessionID = sessionID
	}

	// Send initialized notification
	notification := mcpRequest{
		JSONRPC: "2.0",
		Method:  "notifications/initialized",
		Params:  map[string]interface{}{},
	}

	notifReqBody, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL+"/mcp", bytes.NewReader(notifReqBody))
	if err != nil {
		return fmt.Errorf("failed to create notification request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json, text/event-stream")
	httpReq.Header.Set("X-API-Key", c.apiKey)
	if c.sessionID != "" {
		httpReq.Header.Set("mcp-session-id", c.sessionID)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	defer httpResp.Body.Close()

	c.initialized = true
	return nil
}

// sendRequest sends an MCP request and returns the response with session ID.
func (c *Client) sendRequest(req mcpRequest, sessionID string) (*mcpResponse, string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL+"/mcp", bytes.NewReader(reqBody))
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json, text/event-stream")
	httpReq.Header.Set("X-API-Key", c.apiKey)
	if sessionID != "" {
		httpReq.Header.Set("mcp-session-id", sessionID)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("MCP error (HTTP %d): %s", httpResp.StatusCode, string(body))
	}

	// Check if response is SSE (event-stream)
	contentType := httpResp.Header.Get("Content-Type")
	if contentType == "text/event-stream" {
		// Parse SSE format: event: message\ndata: {...}\n\n
		lines := bytes.Split(body, []byte("\n"))
		var jsonData []byte
		for _, line := range lines {
			if bytes.HasPrefix(line, []byte("data: ")) {
				jsonData = bytes.TrimPrefix(line, []byte("data: "))
				break
			}
		}
		if jsonData == nil {
			return nil, "", fmt.Errorf("no data found in SSE response")
		}
		body = jsonData
	}

	var resp mcpResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, "", fmt.Errorf("failed to parse response (body: %s): %w", string(body), err)
	}

	// Extract session ID from response headers
	responseSessionID := httpResp.Header.Get("mcp-session-id")

	return &resp, responseSessionID, nil
}

// CallTool calls an MCP tool and returns the result.
func (c *Client) CallTool(name string, args map[string]interface{}) (json.RawMessage, error) {
	// Ensure we're initialized
	if err := c.ensureInitialized(); err != nil {
		return nil, err
	}

	// Build tool call request
	req := mcpRequest{
		JSONRPC: "2.0",
		ID:      2, // Use ID 2+ for tool calls (1 was for initialize)
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      name,
			"arguments": args,
		},
	}

	// Send request with session ID
	resp, _, err := c.sendRequest(req, c.sessionID)
	if err != nil {
		return nil, err
	}

	// Check for error
	if resp.Error != nil {
		return nil, fmt.Errorf("tool call error: %s", resp.Error.Message)
	}

	// Parse the result to extract content[0].text
	// MCP returns: {"content": [{"type": "text", "text": "..."}]}
	var toolResult struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(resp.Result, &toolResult); err != nil {
		return nil, fmt.Errorf("failed to parse tool result: %w", err)
	}

	if len(toolResult.Content) == 0 {
		return nil, fmt.Errorf("empty tool result content")
	}

	// Return the text content as raw JSON
	return json.RawMessage(toolResult.Content[0].Text), nil
}
