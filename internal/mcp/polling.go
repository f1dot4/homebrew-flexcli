package mcp

import (
	"encoding/json"
	"fmt"
	"time"
)

// JobStatus represents the status of an async job.
type JobStatus struct {
	Complete        bool            `json:"complete"`
	Status          string          `json:"status"`
	ProgressMessage string          `json:"progress_message"`
	ProgressPercent int             `json:"progress_percent"`
	Result          json.RawMessage `json:"result"`
	Error           string          `json:"error"`
	DownloadURL     string          `json:"download_url"`
	Token           string          `json:"token"`
	Format          string          `json:"format"`
}

// PollForCompletion polls a job until it completes or fails.
// Calls onProgress with new progress messages.
// Returns the final JobStatus when complete.
func PollForCompletion(
	client *Client,
	statusTool string,
	jobID string,
	interval time.Duration,
	onProgress func(string),
) (*JobStatus, error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var lastMessage string

	for {
		// Call status tool
		args := map[string]interface{}{
			"job_id": jobID,
		}

		result, err := client.CallTool(statusTool, args)
		if err != nil {
			return nil, fmt.Errorf("failed to poll job status: %w", err)
		}

		// Parse status
		var status JobStatus
		if err := json.Unmarshal(result, &status); err != nil {
			return nil, fmt.Errorf("failed to parse job status: %w", err)
		}

		// Call progress callback if message changed
		if status.ProgressMessage != "" && status.ProgressMessage != lastMessage {
			if onProgress != nil {
				onProgress(status.ProgressMessage)
			}
			lastMessage = status.ProgressMessage
		}

		// Check if complete
		if status.Complete {
			// Check for error
			if status.Error != "" {
				return &status, fmt.Errorf("job failed: %s", status.Error)
			}
			return &status, nil
		}

		// Wait for next poll
		<-ticker.C
	}
}
