package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// data fitness
// ---------------------------------------------------------------------------

func newDataFitnessCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fitness",
		Short: "View imported fitness data: personal records",
	}

	cmd.AddCommand(newDataFitnessRecordsCmd(rootCfg, resolvedCtx))

	return cmd
}

func newDataFitnessRecordsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "records",
		Short: "List all personal records (PRs) from Garmin",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/fitness/personal-records", nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var records []struct {
				RecordType string  `json:"record_type"`
				Value      float64 `json:"value"`
				Unit       string  `json:"unit"`
				RecordDate string  `json:"record_date"`
				ActivityID string  `json:"activity_id"`
			}
			if err := json.Unmarshal(resp.Data, &records); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			if len(records) == 0 {
				fmt.Println("No personal records found.")
				return nil
			}

			fmt.Println("🏆 Personal Records (Garmin)")
			fmt.Println("==========================")
			fmt.Printf("  %-20s  %-12s  %-12s  %-12s\n", "RECORD TYPE", "VALUE", "DATE", "ACTIVITY ID")
			fmt.Printf("  %-20s  %-12s  %-12s  %-12s\n", "────────────────────", "────────────", "────────────", "────────────")
			for _, r := range records {
				valStr := fmt.Sprintf("%.2f %s", r.Value, r.Unit)
				if r.Unit == "Time" {
					h := int(r.Value) / 3600
					m := (int(r.Value) % 3600) / 60
					s := int(r.Value) % 60
					if h > 0 {
						valStr = fmt.Sprintf("%d:%02d:%02d", h, m, s)
					} else {
						valStr = fmt.Sprintf("%02d:%02d", m, s)
					}
				}
				fmt.Printf("  %-20s  %-12s  %-12s  %-12s\n", r.RecordType, valStr, r.RecordDate, r.ActivityID)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}
