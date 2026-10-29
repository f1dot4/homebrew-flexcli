package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/f1dot4/flexcli/internal/config"
	"github.com/f1dot4/flexcli/internal/mcp"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// data sync
// ---------------------------------------------------------------------------

func newDataSyncCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Manually trigger Garmin or Withings synchronization",
	}

	cmd.AddCommand(newDataSyncGarminCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataSyncWithingsCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataSyncGymCoachCmd(rootCfg, resolvedCtx))

	return cmd
}

func newDataSyncGarminCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "garmin",
		Short: "Sync data from Garmin Connect",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := mcp.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			// Start sync job
			result, err := client.CallTool("sync_garmin", map[string]interface{}{})
			if err != nil {
				return fmt.Errorf("failed to start sync: %w", err)
			}

			// Parse job response
			var jobResp struct {
				JobID  string `json:"job_id"`
				Status string `json:"status"`
			}
			if err := json.Unmarshal(result, &jobResp); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			// Poll for completion
			status, err := mcp.PollForCompletion(
				client,
				"get_sync_status",
				jobResp.JobID,
				1*time.Second,
				func(message string) {
					fmt.Printf("🔄 %s\n", message)
				},
			)
			if err != nil {
				return fmt.Errorf("sync failed: %w", err)
			}

			// Parse result
			var syncResult struct {
				Success          bool `json:"success"`
				ActivitiesSynced int  `json:"activities_synced"`
				WellnessSynced   int  `json:"wellness_synced"`
			}
			if err := json.Unmarshal(status.Result, &syncResult); err == nil {
				fmt.Printf("✅ Garmin synchronization complete: %d activities, %d wellness entries\n",
					syncResult.ActivitiesSynced, syncResult.WellnessSynced)
			} else {
				fmt.Println("✅ Garmin synchronization complete.")
			}

			return nil
		},
	}
}

func newDataSyncWithingsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var startDate string
	var endDate string

	cmd := &cobra.Command{
		Use:   "withings",
		Short: "Sync data from Withings",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateDateFlag(startDate, "--start-date"); err != nil {
				return err
			}
			if err := validateDateFlag(endDate, "--end-date"); err != nil {
				return err
			}
			if (startDate == "") != (endDate == "") {
				return fmt.Errorf("--start-date and --end-date must both be provided, or neither")
			}

			client := mcp.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			// Build arguments
			toolArgs := map[string]interface{}{}
			if startDate != "" {
				toolArgs["start_date"] = startDate
			}
			if endDate != "" {
				toolArgs["end_date"] = endDate
			}

			// Start sync job
			result, err := client.CallTool("sync_withings", toolArgs)
			if err != nil {
				return fmt.Errorf("failed to start sync: %w", err)
			}

			// Parse job response
			var jobResp struct {
				JobID  string `json:"job_id"`
				Status string `json:"status"`
			}
			if err := json.Unmarshal(result, &jobResp); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			// Poll for completion
			status, err := mcp.PollForCompletion(
				client,
				"get_sync_status",
				jobResp.JobID,
				1*time.Second,
				func(message string) {
					fmt.Printf("🔄 %s\n", message)
				},
			)
			if err != nil {
				return fmt.Errorf("sync failed: %w", err)
			}

			// Parse result
			var syncResult struct {
				Success           bool   `json:"success"`
				Message           string `json:"message"`
				MeasurementsSynced int   `json:"measurements_synced"`
			}
			if err := json.Unmarshal(status.Result, &syncResult); err == nil {
				if syncResult.Message != "" {
					fmt.Printf("✅ %s (%d measurements)\n", syncResult.Message, syncResult.MeasurementsSynced)
				} else {
					fmt.Printf("✅ Withings synchronization complete: %d measurements\n", syncResult.MeasurementsSynced)
				}
			} else {
				fmt.Println("✅ Withings synchronization complete.")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&startDate, "start-date", "", "Backfill from this date (YYYY-MM-DD), bypassing the routine sync window; requires --end-date")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Backfill through this date (YYYY-MM-DD), bypassing the routine sync window; requires --start-date")
	return cmd
}
