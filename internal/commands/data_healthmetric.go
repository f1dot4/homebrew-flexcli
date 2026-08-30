package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// data healthmetric (alias: hm)
// ---------------------------------------------------------------------------

func newDataHealthMetricCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "healthmetric",
		Aliases: []string{"hm"},
		Short:   "View imported health metrics (alias: hm): list, show, delete",
	}

	cmd.AddCommand(newDataHealthMetricListCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataHealthMetricShowCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataHealthMetricDeleteCmd(rootCfg, resolvedCtx))

	return cmd
}

func newDataHealthMetricListCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var page int
	var pageSize int
	var asJSON bool
	var startDate string
	var endDate string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List imported health metrics (paginated)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateDateFlag(startDate, "--start-date"); err != nil {
				return err
			}
			if err := validateDateFlag(endDate, "--end-date"); err != nil {
				return err
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := fmt.Sprintf("/api/healthmetrics?page=%d&page_size=%d", page, pageSize)
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
				Metrics []struct {
					ID               int      `json:"id"`
					Date             string   `json:"date"`
					Source           string   `json:"source"`
					WeightKg         *float64 `json:"weight_kg"`
					RestingHeartRate *int     `json:"resting_heart_rate"`
					HRVScore         *float64 `json:"hrv_score"`
					SleepHours       *float64 `json:"sleep_hours"`
					CyclingFTP       *float64 `json:"cycling_ftp"`
					CyclingLTHR      *float64 `json:"cycling_lthr"`
					RunningFTP       *float64 `json:"running_ftp"`
					RunningLTHR      *float64 `json:"running_lthr"`
					ActiveCalories   *int     `json:"active_calories"`
					PassiveCalories  *int     `json:"passive_calories"`
					TotalCalories    *int     `json:"total_calories"`
				} `json:"metrics"`
				TotalEntries int `json:"total_entries"`
				TotalPages   int `json:"total_pages"`
				CurrentPage  int `json:"current_page"`
			}
			if err := json.Unmarshal(resp.Data, &data); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			if len(data.Metrics) == 0 {
				fmt.Println("No health metrics found.")
				return nil
			}

			fmt.Printf("Health metrics (page %d/%d, %d total):\n\n", data.CurrentPage, data.TotalPages, data.TotalEntries)
			fmt.Printf("  %-12s  %-10s  %8s  %5s  %5s  %5s  %6s  %6s  %6s  %6s  %6s  %6s  %6s\n",
				"DATE", "SOURCE", "WEIGHT", "RHR", "HRV", "SLEEP", "C-FTP", "C-LTHR", "R-FTP", "R-LTHR", "ACAL", "PCAL", "TCAL")
			fmt.Printf("  %-12s  %-10s  %8s  %5s  %5s  %5s  %6s  %6s  %6s  %6s  %6s  %6s  %6s\n",
				"──────────", "──────────", "────────", "─────", "─────", "─────", "──────", "──────", "──────", "──────", "──────", "──────", "──────")
			for _, m := range data.Metrics {
				weight := "  -"
				if m.WeightKg != nil {
					weight = fmt.Sprintf("%6.1fkg", *m.WeightKg)
				}
				rhr := "  -"
				if m.RestingHeartRate != nil {
					rhr = fmt.Sprintf("%5d", *m.RestingHeartRate)
				}
				hrv := "  -"
				if m.HRVScore != nil {
					hrv = fmt.Sprintf("%5.1f", *m.HRVScore)
				}
				sleep := "  -"
				if m.SleepHours != nil {
					sleep = fmt.Sprintf("%5.1f", *m.SleepHours)
				}
				cftp := "  -"
				if m.CyclingFTP != nil {
					cftp = fmt.Sprintf("%6.0f", *m.CyclingFTP)
				}
				clthr := "  -"
				if m.CyclingLTHR != nil {
					clthr = fmt.Sprintf("%6.0f", *m.CyclingLTHR)
				}
				rftp := "  -"
				if m.RunningFTP != nil {
					rftp = fmt.Sprintf("%6.0f", *m.RunningFTP)
				}
				rlthr := "  -"
				if m.RunningLTHR != nil {
					rlthr = fmt.Sprintf("%6.0f", *m.RunningLTHR)
				}
				acal := "  -"
				if m.ActiveCalories != nil {
					acal = fmt.Sprintf("%6d", *m.ActiveCalories)
				}
				pcal := "  -"
				if m.PassiveCalories != nil {
					pcal = fmt.Sprintf("%6d", *m.PassiveCalories)
				}
				tcal := "  -"
				if m.TotalCalories != nil {
					tcal = fmt.Sprintf("%6d", *m.TotalCalories)
				}
				fmt.Printf("  %-12s  %-10s  %8s  %5s  %5s  %5s  %6s  %6s  %6s  %6s  %6s  %6s  %6s\n",
					m.Date, m.Source, weight, rhr, hrv, sleep, cftp, clthr, rftp, rlthr, acal, pcal, tcal)
			}

			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 20, "Number of metrics per page")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Only include metrics on or after this date (YYYY-MM-DD)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Only include metrics on or before this date (YYYY-MM-DD)")
	return cmd
}

func newDataHealthMetricShowCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "show [date]",
		Short: "Show aggregated health metric for a specific date (YYYY-MM-DD, defaults to today)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			date := time.Now().Format("2006-01-02")
			if len(args) > 0 {
				date = args[0]
			}
			if date == "today" {
				date = time.Now().Format("2006-01-02")
			}
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/healthmetric/"+date, nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var metric map[string]interface{}
			if err := json.Unmarshal(resp.Data, &metric); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			fmt.Printf("Health metric for %s\n", date)
			for k, v := range metric {
				if v == nil {
					continue
				}
				fmt.Printf("  • %-28s %v\n", k+":", v)
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newDataHealthMetricDeleteCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a health metric (currently disabled)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Health metric deletion is currently not implemented due to safety reasons.")
			fmt.Println("Please remove records manually if absolutely necessary.")
			return nil
		},
	}
}
