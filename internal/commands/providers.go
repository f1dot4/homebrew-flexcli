package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewGarminConnectionCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "garmin",
		Short: "Manage Garmin connection and settings",
	}

	cmd.AddCommand(NewProviderConfigCmd("garmin", rootCfg, resolvedCtx))
	return cmd
}

func NewWithingsConnectionCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withings",
		Short: "Manage Withings connection and settings",
	}

	cmd.AddCommand(NewProviderConfigCmd("withings", rootCfg, resolvedCtx))
	return cmd
}

func NewProviderConfigCmd(provider string, rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: fmt.Sprintf("Manage %s expert settings", provider),
	}

	cmd.AddCommand(newConfigGetCmd(provider, resolvedCtx))
	cmd.AddCommand(newConfigSetCmd(provider, resolvedCtx))

	return cmd
}

func newConfigGetCmd(provider string, ctx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get expert settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(ctx.ServerURL, ctx.APIKey)
			resp, err := client.Request("GET", "/api/profile/expert-settings", nil)
			if err != nil {
				return err
			}

			var settings map[string]interface{}
			if err := json.Unmarshal(resp.Data, &settings); err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			if provider == "garmin" {
				fmt.Fprintln(w, "SETTING\tVALUE")
				fmt.Fprintf(w, "Sync Interval (Hours)\t%v\n", settings["garmin_sync_interval_hours"])
				fmt.Fprintf(w, "Manual Sync Lookback (Days)\t%v\n", settings["sync_days_manual"])
				fmt.Fprintf(w, "Scheduled Sync Lookback (Days)\t%v\n", settings["sync_days_schedule"])
			} else if provider == "withings" {
				fmt.Fprintln(w, "SETTING\tVALUE")
				fmt.Fprintf(w, "Sync Interval (Hours)\t%v\n", settings["withings_sync_interval_hours"])
			}
			w.Flush()
			return nil
		},
	}
}

func newConfigSetCmd(provider string, ctx *config.Context) *cobra.Command {
	var interval int
	var lookbackManual int
	var lookbackSchedule int

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Update expert settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			updates := make(map[string]interface{})

			if provider == "garmin" {
				if cmd.Flags().Changed("interval") {
					updates["garmin_sync_interval_hours"] = interval
				}
			} else if provider == "withings" {
				if cmd.Flags().Changed("interval") {
					updates["withings_sync_interval_hours"] = interval
				}
			}

			// Universal settings (apply to both or specific?)
			// The plan implies these are provider specific in the command, but the model shares them.
			// Let's allow setting them under Garmin for now as they are most relevant there.
			if provider == "garmin" {
				if cmd.Flags().Changed("lookback-manual") {
					updates["sync_days_manual"] = lookbackManual
				}
				if cmd.Flags().Changed("lookback-schedule") {
					updates["sync_days_schedule"] = lookbackSchedule
				}
			}

			if len(updates) == 0 {
				return fmt.Errorf("no settings provided to update")
			}

			client := api.NewClient(ctx.ServerURL, ctx.APIKey)
			_, err := client.Request("POST", "/api/profile/expert-settings", updates)
			if err != nil {
				return err
			}

			fmt.Println("✅ Expert settings updated successfully.")
			return nil
		},
	}

	cmd.Flags().IntVar(&interval, "interval", 0, "Sync interval in hours")
	if provider == "garmin" {
		cmd.Flags().IntVar(&lookbackManual, "lookback-manual", 0, "Days to look back for manual sync")
		cmd.Flags().IntVar(&lookbackSchedule, "lookback-schedule", 0, "Days to look back for scheduled sync")
	}

	return cmd
}
