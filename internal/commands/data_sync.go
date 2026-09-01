package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
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

	return cmd
}

func newDataSyncGarminCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "garmin",
		Short: "Sync data from Garmin Connect",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			events, err := client.PostSSE("/api/sync/garmin/stream")
			if err != nil {
				return err
			}

			var syncSuccess bool
			var finalError string

			for event := range events {
				if event.Event == "result" {
					var result struct {
						Success bool   `json:"success"`
						Error   string `json:"error"`
					}
					if err := json.Unmarshal([]byte(event.Data), &result); err == nil {
						syncSuccess = result.Success
						finalError = result.Error
					}
				} else {
					fmt.Printf("🔄 %s\n", event.Data)
				}
			}

			if syncSuccess {
				fmt.Println("✅ Garmin synchronization complete.")
			} else {
				if finalError != "" {
					fmt.Printf("❌ Garmin synchronization failed: %s\n", finalError)
				} else {
					fmt.Println("❌ Garmin synchronization failed.")
				}
				return fmt.Errorf("sync failed")
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

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := "/api/sync/withings/stream"
			if dateQuery := buildDateRangeQuery(startDate, endDate); dateQuery != "" {
				path += "?" + dateQuery
			}
			events, err := client.PostSSE(path)
			if err != nil {
				return err
			}

			var syncSuccess bool
			var finalError string

			for event := range events {
				if event.Event == "result" {
					var result struct {
						Success bool   `json:"success"`
						Error   string `json:"error"`
					}
					if err := json.Unmarshal([]byte(event.Data), &result); err == nil {
						syncSuccess = result.Success
						finalError = result.Error
					}
				} else {
					fmt.Printf("🔄 %s\n", event.Data)
				}
			}

			if syncSuccess {
				fmt.Println("✅ Withings synchronization complete.")
			} else {
				if finalError != "" {
					fmt.Printf("❌ Withings synchronization failed: %s\n", finalError)
				} else {
					fmt.Println("❌ Withings synchronization failed.")
				}
				return fmt.Errorf("sync failed")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&startDate, "start-date", "", "Backfill from this date (YYYY-MM-DD), bypassing the routine sync window; requires --end-date")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Backfill through this date (YYYY-MM-DD), bypassing the routine sync window; requires --start-date")
	return cmd
}
