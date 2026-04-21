package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func newPreferencesCustomCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "custom",
		Short: "Manage custom training preferences (free-text list)",
	}

	cmd.AddCommand(newPreferencesCustomListCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newPreferencesCustomAddCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newPreferencesCustomRemoveCmd(rootCfg, resolvedCtx))

	return cmd
}

func newPreferencesCustomListCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List custom training preferences",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/profile/preferences/custom", nil)
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

			preferences := data["preferences"].([]interface{})
			fmt.Println("📝 Custom Training Preferences:")
			if len(preferences) == 0 {
				fmt.Println("  None")
			}
			for i, p := range preferences {
				fmt.Printf("  %d. %s\n", i, p)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newPreferencesCustomAddCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "add [text]",
		Short: "Add a new custom training preference",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			payload := map[string]string{
				"text": args[0],
			}

			resp, err := client.Request("POST", "/api/profile/preferences/custom", payload)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			fmt.Println("✅ Preference added successfully.")
			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newPreferencesCustomRemoveCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "remove [index]",
		Short: "Remove a custom training preference by index",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := fmt.Sprintf("/api/profile/preferences/custom/%s", args[0])
			resp, err := client.Request("DELETE", path, nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Printf("{\"success\": true, \"message\": \"%s\"}\n", resp.Message)
				return nil
			}

			fmt.Println(resp.Message)
			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}
