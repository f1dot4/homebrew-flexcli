package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// data activity (alias: act)
// ---------------------------------------------------------------------------

func newDataActivityCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "activity",
		Aliases: []string{"act"},
		Short:   "Manage Garmin activities (alias: act): list, download, upload, delete",
	}

	cmd.AddCommand(newDataActivityListCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataActivityDownloadCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataActivityDownloadBulkCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataActivityUploadCmd(rootCfg, resolvedCtx))

	cmd.AddCommand(newDataActivityDeleteCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataActivityRenameCmd(rootCfg, resolvedCtx))

	return cmd
}

func newDataActivityListCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var page int
	var pageSize int
	var asJSON bool
	var startDate string
	var endDate string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List synced activities with their Garmin activity IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateDateFlag(startDate, "--start-date"); err != nil {
				return err
			}
			if err := validateDateFlag(endDate, "--end-date"); err != nil {
				return err
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := fmt.Sprintf("/api/activities?page=%d&page_size=%d", page, pageSize)
			if dateQuery := buildDateRangeQuery(startDate, endDate); dateQuery != "" {
				path += "&" + dateQuery
			}
			resp, err := client.Request("GET", path, nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var data struct {
				Activities []struct {
					GarminActivityID string   `json:"garmin_activity_id"`
					Type             string   `json:"type"`
					Description      string   `json:"description"`
					StartTime        string   `json:"start_time"`
					DurationMinutes  int      `json:"duration_minutes"`
					DistanceKm       *float64 `json:"distance_km"`
				} `json:"activities"`
				TotalEntries int `json:"total_entries"`
				TotalPages   int `json:"total_pages"`
				CurrentPage  int `json:"current_page"`
			}
			if err := json.Unmarshal(resp.Data, &data); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			if len(data.Activities) == 0 {
				fmt.Println("No activities found.")
				return nil
			}

			fmt.Printf("Activities (page %d/%d, %d total):\n\n", data.CurrentPage, data.TotalPages, data.TotalEntries)
			fmt.Printf("  %-16s  %-20s  %-22s  %6s  %s\n", "GARMIN ID", "DATE", "TYPE", "MIN", "DESCRIPTION")
			fmt.Printf("  %-16s  %-20s  %-22s  %6s  %s\n", "────────────────", "────────────────────", "──────────────────────", "──────", "───────────")
			for _, a := range data.Activities {
				dateStr := a.StartTime
				if len(dateStr) > 16 {
					dateStr = dateStr[:16]
				}
				actType := strings.Replace(a.Type, "_", " ", -1)
				desc := a.Description
				fmt.Printf("  %-16s  %-20s  %-22s  %6d  %s\n", a.GarminActivityID, dateStr, actType, a.DurationMinutes, desc)
			}

			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 20, "Number of activities per page")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Only include activities on or after this date (YYYY-MM-DD)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Only include activities on or before this date (YYYY-MM-DD)")
	return cmd
}

func newDataActivityDownloadCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var output string
	var format string

	cmd := &cobra.Command{
		Use:   "download [activity_id]",
		Short: "Download an activity file from Garmin Connect (defaults to 'latest')",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			activityID := "latest"
			if len(args) > 0 {
				activityID = args[0]
			}
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			if output == "" {
				ext := format
				if format == "fit" {
					ext = "zip"
				}
				output = activityID + "." + ext
			}

			path := fmt.Sprintf("/api/activity/%s/download?format=%s", activityID, format)
			if err := client.DownloadFile(path, output); err != nil {
				return fmt.Errorf("download failed: %w", err)
			}

			fmt.Printf("Downloaded activity %s to %s\n", activityID, output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file path (default: <activity_id>.<ext>)")
	cmd.Flags().StringVarP(&format, "format", "f", "fit", "File format (fit, gpx, tcx, csv, kml)")
	return cmd
}

func newDataActivityDownloadBulkCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var output string
	var format string
	var year int
	var month int
	var day int

	cmd := &cobra.Command{
		Use:   "download-bulk",
		Short: "Download multiple activities in a ZIP bundle",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			if year == 0 {
				return fmt.Errorf("year is required")
			}

			path := fmt.Sprintf("/api/activity/download/bulk/prepare?year=%d&format=%s", year, format)
			if month > 0 {
				path += fmt.Sprintf("&month=%d", month)
			}
			if day > 0 {
				path += fmt.Sprintf("&day=%d", day)
			}

			events, err := client.GetSSE(path)
			if err != nil {
				return err
			}

			var token string
			var finalErr error

			for event := range events {
				var data map[string]interface{}
				if err := json.Unmarshal([]byte(event.Data), &data); err != nil {
					continue
				}

				status, _ := data["status"].(string)
				switch status {
				case "progress":
					msg, _ := data["message"].(string)
					fmt.Printf("📦 %s\n", msg)
				case "complete":
					token, _ = data["token"].(string)
					fmt.Println("✅ Bulk preparation complete.")
				case "error":
					msg, _ := data["message"].(string)
					finalErr = fmt.Errorf("bulk preparation failed: %s", msg)
				}
			}

			if finalErr != nil {
				return finalErr
			}

			if token == "" {
				return fmt.Errorf("no download token received")
			}

			if output == "" {
				output = fmt.Sprintf("activities_%d.zip", year)
				if month > 0 {
					output = fmt.Sprintf("activities_%d_%02d.zip", year, month)
					if day > 0 {
						output = fmt.Sprintf("activities_%d_%02d_%02d.zip", year, month, day)
					}
				}
			}

			fmt.Printf("Downloading bundle to %s...\n", output)
			downloadPath := fmt.Sprintf("/api/activity/download/bulk/%s", token)
			if err := client.DownloadFile(downloadPath, output); err != nil {
				return fmt.Errorf("download failed: %w", err)
			}

			fmt.Printf("✅ Bulk download saved to %s\n", output)
			return nil
		},
	}

	cmd.Flags().IntVarP(&year, "year", "y", 0, "Year to download activities for (required)")
	cmd.Flags().IntVarP(&month, "month", "m", 0, "Month to download (1-12, optional)")
	cmd.Flags().IntVarP(&day, "day", "d", 0, "Day to download (1-31, optional)")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output ZIP file path")
	cmd.Flags().StringVarP(&format, "format", "f", "fit", "File format (fit, gpx, tcx, csv, kml)")

	return cmd
}

func newDataActivityUploadCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "upload <file>",
		Short: "Upload a FIT/GPX/TCX file to Garmin Connect",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return fmt.Errorf("file not found: %s", filePath)
			}
			ext := strings.ToLower(filepath.Ext(filePath))
			if ext != ".fit" && ext != ".gpx" && ext != ".tcx" {
				return fmt.Errorf("unsupported file type %s (must be .fit, .gpx, or .tcx)", ext)
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.UploadFile("/api/activity/upload", filePath)
			if err != nil {
				return fmt.Errorf("upload failed: %w", err)
			}

			if resp.Success {
				fmt.Printf("Uploaded %s successfully.\n", filepath.Base(filePath))
			} else {
				fmt.Printf("Upload failed: %s\n", resp.Message)
			}
			return nil
		},
	}
}

func newDataActivityDeleteCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [activity_id]",
		Short: "Delete an activity from Garmin Connect (defaults to 'latest', currently disabled)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Activity deletion is currently not implemented due to safety reasons.")
			fmt.Println("Please delete activities manually via Garmin Connect.")
			return nil
		},
	}
}

func newDataActivityRenameCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "rename <title> [activity_id]",
		Short: "Rename an activity in Garmin Connect (defaults to 'latest')",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			title := args[0]
			activityID := "latest"
			if len(args) > 1 {
				activityID = args[1]
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			body := map[string]string{
				"name": title,
			}

			path := fmt.Sprintf("/api/activity/%s/name", activityID)
			resp, err := client.Request("PUT", path, body)
			if err != nil {
				return fmt.Errorf("rename failed: %w", err)
			}

			if resp.Success {
				fmt.Printf("✅ Activity %s renamed to: %s\n", activityID, title)
			} else {
				fmt.Printf("❌ Rename failed: %s\n", resp.Message)
			}
			return nil
		},
	}
}
