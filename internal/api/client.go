package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Client struct {
	BaseURL string
	APIKey  string
	IsAdmin bool
	HTTP    *http.Client
}

type APIResponse struct {
	Success    bool            `json:"success"`
	Status     string          `json:"status,omitempty"`
	Message    string          `json:"message,omitempty"`
	Detail     string          `json:"detail,omitempty"` // FastAPI HTTPException format
	Data       json.RawMessage `json:"data,omitempty"`
	Plan       json.RawMessage `json:"plan,omitempty"`       // Legacy, keeping for now
	Thresholds json.RawMessage `json:"thresholds,omitempty"` // Legacy
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: strings.TrimSuffix(baseURL, "/"),
		APIKey:  apiKey,
		HTTP:    &http.Client{},
	}
}

func (c *Client) Request(method, path string, body interface{}) (*APIResponse, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if c.IsAdmin {
		req.Header.Set("X-Admin-Key", c.APIKey)
	} else {
		req.Header.Set("X-API-Key", c.APIKey)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: check your API key")
	}

	if resp.StatusCode >= 400 {
		var errResp APIResponse
		json.NewDecoder(resp.Body).Decode(&errResp)
		if errResp.Message != "" {
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp.Message)
		}
		if errResp.Detail != "" {
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp.Detail)
		}
		return nil, fmt.Errorf("API error (%d)", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// DownloadFile performs a GET request and saves the response body to destPath.
func (c *Client) DownloadFile(path string, destPath string) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if c.IsAdmin {
		req.Header.Set("X-Admin-Key", c.APIKey)
	} else {
		req.Header.Set("X-API-Key", c.APIKey)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("unauthorized: check your API key")
	}
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", destPath, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// UploadFile performs a multipart POST, uploading the file at filePath.
func (c *Client) UploadFile(path string, filePath string) (*APIResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}
	writer.Close()

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}

	if c.IsAdmin {
		req.Header.Set("X-Admin-Key", c.APIKey)
	} else {
		req.Header.Set("X-API-Key", c.APIKey)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: check your API key")
	}
	if resp.StatusCode >= 400 {
		var errResp APIResponse
		json.NewDecoder(resp.Body).Decode(&errResp)
		if errResp.Message != "" {
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp.Message)
		}
		if errResp.Detail != "" {
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp.Detail)
		}
		return nil, fmt.Errorf("API error (%d)", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &apiResp, nil
}
