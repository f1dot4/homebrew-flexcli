package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewGoalCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goal",
		Short: "Manage training goals",
	}

	cmd.AddCommand(newGoalListCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGoalAddCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGoalDeleteCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGoalSuggestCmd(rootCfg, resolvedCtx))

	return cmd
}

func newGoalListCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List active and pending goals",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			resp, err := client.Request("GET", "/api/goals", nil)
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

			active := data["active"].([]interface{})
			pending := data["pending"].([]interface{})

			fmt.Println("🎯 Active Goals:")
			if len(active) == 0 {
				fmt.Println("  None")
			}
			for _, g := range active {
				goal := g.(map[string]interface{})
				fmt.Printf("  • %s (ID: %s)\n", goal["name"], goal["goal_id"])
			}

			fmt.Println("\n⏳ Pending Goals:")
			if len(pending) == 0 {
				fmt.Println("  None")
			}
			for _, g := range pending {
				goal := g.(map[string]interface{})
				fmt.Printf("  • %s (ID: %s)\n", goal["name"], goal["goal_id"])
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newGoalAddCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var description string
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add a new performance goal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			payload := map[string]string{
				"name":        args[0],
				"description": description,
			}

			resp, err := client.Request("POST", "/api/goals", payload)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			fmt.Println("✅ Goal added successfully.")
			return nil
		},
	}

	cmd.Flags().StringVar(&description, "description", "", "Goal description")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}

func newGoalDeleteCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a goal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := fmt.Sprintf("/api/goals/%s", args[0])
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

func newGoalSuggestCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "suggest [goal description]",
		Short: "Suggest measurable targets for a qualitative goal using AI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			payload := map[string]string{
				"goal": args[0],
			}

			resp, err := client.Request("POST", "/api/goals/suggest", payload)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var payloadData map[string]interface{}
			if err := json.Unmarshal(resp.Data, &payloadData); err != nil {
				return err
			}

			targets, ok := payloadData["targets"].([]interface{})
			if !ok {
				return fmt.Errorf("unexpected response format: missing 'targets' field")
			}

			fmt.Println("🤖 AI Goal Suggestions:")
			if len(targets) == 0 {
				fmt.Println("  No suggestions found.")
			}
			for i, t := range targets {
				target, ok := t.(map[string]interface{})
				if !ok {
					continue
				}
				fmt.Printf("\n  Target %d:\n", i+1)
				fmt.Printf("  • Metric:    %s\n", target["metric"])
				fmt.Printf("  • Value:     %v %s\n", target["value"], target["unit"])
				fmt.Printf("  • Operator:  %s\n", target["operator"])
				fmt.Printf("  • Reasoning: %s\n", target["reasoning"])
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output in JSON format")
	return cmd
}
