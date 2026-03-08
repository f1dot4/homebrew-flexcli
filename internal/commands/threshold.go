package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewThresholdCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "threshold",
		Short: "Manage training thresholds",
	}

	cmd.AddCommand(newThresholdGetCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newThresholdSetCmd(rootCfg, resolvedCtx))

	return cmd
}

func newThresholdGetCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get current thresholds",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/thresholds", nil)
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

			if data["thresholds"] == nil {
				fmt.Println("No thresholds found.")
				return nil
			}

			thresholds := data["thresholds"].(map[string]interface{})
			fmt.Println("📊 Current Training Thresholds")
			fmt.Println("------------------------------")

			// Helper to get status hint
			getHint := func(isLearned, isDerived interface{}) string {
				if il, ok := isLearned.(bool); ok && il {
					return " 📈"
				}
				if id, ok := isDerived.(bool); ok && id {
					return " 🔢"
				}
				return ""
			}

			fmt.Println("Running:")
			fmt.Printf("  • FTP:   %v W%s\n", thresholds["effective_running_ftp"], getHint(thresholds["is_running_ftp_learned"], thresholds["is_running_ftp_derived"]))
			fmt.Printf("  • LTHR:  %v bpm%s\n", thresholds["effective_running_lthr"], getHint(thresholds["is_running_lthr_learned"], thresholds["is_running_lthr_derived"]))
			fmt.Printf("  • Pace:  %v%s\n", thresholds["effective_running_threshold_pace"], getHint(thresholds["is_running_pace_learned"], nil))

			fmt.Println("\nCycling:")
			fmt.Printf("  • FTP:   %v W%s\n", thresholds["effective_cycling_ftp"], getHint(thresholds["is_cycling_ftp_learned"], thresholds["is_cycling_ftp_derived"]))
			fmt.Printf("  • LTHR:  %v bpm%s\n", thresholds["effective_cycling_lthr"], getHint(thresholds["is_cycling_lthr_learned"], thresholds["is_cycling_lthr_derived"]))
			fmt.Printf("  • Pace:  %v%s\n", thresholds["effective_cycling_threshold_pace"], getHint(thresholds["is_cycling_pace_learned"], thresholds["is_cycling_pace_derived"]))

			hasLearned := false
			hasDerived := false
			for k, v := range thresholds {
				if k[len(k)-7:] == "learned" {
					if b, ok := v.(bool); ok && b {
						hasLearned = true
					}
				}
				if k[len(k)-7:] == "derived" {
					if b, ok := v.(bool); ok && b {
						hasDerived = true
					}
				}
			}

			if hasLearned || hasDerived {
				fmt.Println("")
				if hasLearned {
					fmt.Println("📈 = learned from history")
				}
				if hasDerived {
					fmt.Println("🔢 = calculated via formula")
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newThresholdSetCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var rFTP, rLTHR, cFTP, cLTHR int
	var rPace, cPace string
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set thresholds",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			update := make(map[string]interface{})
			if cmd.Flags().Changed("running-ftp") {
				update["running_ftp"] = rFTP
			}
			if cmd.Flags().Changed("running-lthr") {
				update["running_lthr"] = rLTHR
			}
			if cmd.Flags().Changed("running-pace") {
				update["running_threshold_pace"] = rPace
			}
			if cmd.Flags().Changed("cycling-ftp") {
				update["cycling_ftp"] = cFTP
			}
			if cmd.Flags().Changed("cycling-lthr") {
				update["cycling_lthr"] = cLTHR
			}
			if cmd.Flags().Changed("cycling-pace") {
				update["cycling_threshold_pace"] = cPace
			}

			if len(update) == 0 {
				return fmt.Errorf("no threshold values provided to set")
			}

			resp, err := client.Request("POST", "/api/thresholds", update)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			fmt.Println("✅ Thresholds updated successfully.")
			return nil
		},
	}

	cmd.Flags().IntVar(&rFTP, "running-ftp", 0, "Running FTP (W)")
	cmd.Flags().IntVar(&rLTHR, "running-lthr", 0, "Running LTHR (bpm)")
	cmd.Flags().StringVar(&rPace, "running-pace", "", "Running Pace (e.g. 4:30/km)")
	cmd.Flags().IntVar(&cFTP, "cycling-ftp", 0, "Cycling FTP (W)")
	cmd.Flags().IntVar(&cLTHR, "cycling-lthr", 0, "Cycling LTHR (bpm)")
	cmd.Flags().StringVar(&cPace, "cycling-pace", "", "Cycling Pace (e.g. 1:20/km)")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")

	return cmd
}
