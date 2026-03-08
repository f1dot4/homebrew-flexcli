package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewStatsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "View training statistics and reports",
	}

	cmd.AddCommand(newStatsDashboardCmd(rootCfg, resolvedCtx))

	return cmd
}

func newStatsDashboardCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "View training dashboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/stats/dashboard", nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var data map[string]interface{}
			if err := json.Unmarshal(resp.Data, &data); err != nil {
				return err
			}

			fmt.Println("📊 Training Dashboard")
			fmt.Println("======================")

			phys, _ := data["physiological_status"].(map[string]interface{})
			if phys != nil {
				fmt.Printf("\n🧬 Physiological Status:\n")
				fmt.Printf("  • Form:    %v %v (TSB: %v)\n", phys["emoji"], phys["label"], phys["tsb"])
				fmt.Printf("  • Fitness: %v (CTL) | Fatigue: %v (ATL)\n", phys["ctl"], phys["atl"])
			}

			adherence, _ := data["adherence"].(map[string]interface{})
			if adherence != nil {
				fmt.Printf("\n📋 Plan Adherence:\n")
				fmt.Printf("  • %v%% (%v/%v sessions)\n",
					adherence["adherence_percentage"],
					adherence["completed_count"],
					adherence["planned_count"])
			}

			trends, _ := data["vital_trends"].([]interface{})
			if len(trends) > 0 {
				fmt.Printf("\n❤️ Vital Trends:\n")
				for _, t := range trends {
					trend := t.(map[string]interface{})
					fmt.Printf("  • %s: %v %s (%s)\n", trend["label"], trend["current"], trend["unit"], trend["trend"])
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}
