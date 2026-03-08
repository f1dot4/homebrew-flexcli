package commands

import (
	"fmt"
	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewSyncCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Manually synchronize health and activity data",
	}

	cmd.AddCommand(NewSyncGarminCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(NewSyncWithingsCmd(rootCfg, resolvedCtx))

	return cmd
}

func NewSyncGarminCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "garmin",
		Short: "Sync data from Garmin Connect",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("POST", "/api/sync/garmin", nil)
			if err != nil {
				return err
			}

			if resp.Success {
				fmt.Println("✅ Garmin synchronization triggered successfully.")
				if resp.Message != "" && resp.Message != "Garmin sync triggered" {
					fmt.Printf("Details: %s\n", resp.Message)
				}
			} else {
				fmt.Printf("❌ Garmin synchronization failed: %s\n", resp.Message)
			}

			return nil
		},
	}
}

func NewSyncWithingsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "withings",
		Short: "Sync data from Withings",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("POST", "/api/sync/withings", nil)
			if err != nil {
				return err
			}

			if resp.Success {
				fmt.Println("✅ Withings synchronization triggered successfully.")
				if resp.Message != "" && resp.Message != "Withings sync triggered" {
					fmt.Printf("Details: %s\n", resp.Message)
				}
			} else {
				fmt.Printf("❌ Withings synchronization failed: %s\n", resp.Message)
			}

			return nil
		},
	}
}
