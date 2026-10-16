package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// ---------------------------------------------------------------------------
// profile connect gymcoach
// ---------------------------------------------------------------------------

func NewGymCoachConnectionCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gymcoach",
		Short: "Manage GymCoach connection and settings",
	}

	cmd.AddCommand(newGymCoachConfigCmd(rootCfg, resolvedCtx))
	return cmd
}

func newGymCoachConfigCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage the GymCoach server URL and MCP token",
	}

	cmd.AddCommand(newGymCoachConfigGetCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGymCoachConfigSetCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGymCoachConfigDeleteCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGymCoachConfigTestCmd(rootCfg, resolvedCtx))

	return cmd
}

func newGymCoachConfigGetCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Show the current GymCoach configuration and status",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/connect/gymcoach/status", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var status struct {
				Connected bool    `json:"connected"`
				BaseURL   string  `json:"base_url"`
				LastSync  *string `json:"last_sync"`
			}
			if err := json.Unmarshal(resp.Data, &status); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}

			connected := "❌ Not connected"
			if status.Connected {
				connected = "✅ Connected"
			}
			lastSync := "-"
			if status.LastSync != nil {
				lastSync = *status.LastSync
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "SETTING\tVALUE")
			fmt.Fprintf(w, "Status\t%s\n", connected)
			fmt.Fprintf(w, "Server URL\t%s\n", status.BaseURL)
			fmt.Fprintf(w, "Last Sync\t%s\n", lastSync)
			w.Flush()
			return nil
		},
	}

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGymCoachConfigSetCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var url string
	var token string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set the GymCoach server URL and/or MCP token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if url == "" && token == "" {
				return fmt.Errorf("nothing to update: provide --url and/or --token")
			}

			if token == "" {
				fmt.Print("Enter GymCoach MCP token: ")
				byteToken, err := term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					return err
				}
				fmt.Println()
				token = strings.TrimSpace(string(byteToken))
			}

			body := map[string]interface{}{}
			if url != "" {
				body["base_url"] = url
			}
			if token != "" {
				body["mcp_token"] = token
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("POST", "/api/connect/gymcoach/config", body)
			if err != nil {
				return err
			}
			if !resp.Success {
				fmt.Printf("❌ %s\n", resp.Message)
				return fmt.Errorf("failed to save configuration")
			}

			fmt.Printf("✅ %s\n", resp.Message)
			return nil
		},
	}

	cmd.Flags().StringVar(&url, "url", "", "GymCoach server URL (optional; blank keeps the default)")
	cmd.Flags().StringVar(&token, "token", "", "GymCoach MCP token (prompted if omitted)")
	return cmd
}

func newGymCoachConfigDeleteCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete the stored GymCoach MCP token",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("DELETE", "/api/connect/gymcoach/config", nil)
			if err != nil {
				return err
			}
			if !resp.Success {
				fmt.Printf("❌ %s\n", resp.Message)
				return fmt.Errorf("failed to delete configuration")
			}
			fmt.Printf("✅ %s\n", resp.Message)
			return nil
		},
	}
}

func newGymCoachConfigTestCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test the GymCoach connection by listing today's trainings",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/connect/gymcoach/trainings?days=1", nil)
			if err != nil {
				return err
			}
			if !resp.Success {
				fmt.Printf("❌ Connection failed: %s\n", resp.Message)
				return fmt.Errorf("connection test failed")
			}
			fmt.Println("✅ Connection OK: GymCoach API responded.")
			return nil
		},
	}
}

// ---------------------------------------------------------------------------
// profile data gymcoach
// ---------------------------------------------------------------------------

func NewGymCoachDataCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gymcoach",
		Short: "Inspect GymCoach trainings",
	}

	cmd.AddCommand(newGymCoachTrainingsCmd(rootCfg, resolvedCtx))
	return cmd
}

func newGymCoachTrainingsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var days int
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "trainings",
		Short: "List recent GymCoach trainings",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request(
				"GET",
				fmt.Sprintf("/api/connect/gymcoach/trainings?days=%d", days),
				nil,
			)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var data struct {
				Count     int `json:"count"`
				Trainings []struct {
					Date            string  `json:"date"`
					WorkoutName     string  `json:"workout_name"`
					ExerciseCount   int     `json:"exercise_count"`
					SetCount        int     `json:"set_count"`
					TotalVolumeKg   float64 `json:"total_volume_kg"`
					DurationMinutes *int    `json:"duration_minutes"`
				} `json:"trainings"`
			}
			if err := json.Unmarshal(resp.Data, &data); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if data.Count == 0 {
				fmt.Printf("No GymCoach trainings found in the last %d days.\n", days)
				return nil
			}

			fmt.Printf("🏋️  GymCoach trainings (last %d days)\n", days)
			fmt.Println("=====================================")
			fmt.Printf("  %-12s  %-25s  %-10s  %-8s  %-12s\n", "DATE", "WORKOUT", "EXERCISES", "SETS", "VOLUME")
			fmt.Printf("  %-12s  %-25s  %-10s  %-8s  %-12s\n", "────────────", "─────────────────────────", "──────────", "────────", "────────────")
			for _, t := range data.Trainings {
				fmt.Printf("  %-12s  %-25s  %-10d  %-8d  %-12s\n",
					t.Date, t.WorkoutName, t.ExerciseCount, t.SetCount,
					fmt.Sprintf("%.0f kg", t.TotalVolumeKg))
			}
			return nil
		},
	}

	cmd.Flags().IntVar(&days, "days", 30, "Number of days to look back")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

// ---------------------------------------------------------------------------
// profile data sync gymcoach
// ---------------------------------------------------------------------------

func newDataSyncGymCoachCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var trainings bool
	var weight bool
	var startDate string
	var endDate string

	cmd := &cobra.Command{
		Use:   "gymcoach",
		Short: "Sync GymCoach: push trainings to Garmin and set bodyweight",
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

			// Default to doing both when no explicit selection is made.
			if !trainings && !weight {
				trainings = true
				weight = true
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			if trainings {
				path := "/api/activity/import/gymcoach"
				if dateQuery := buildDateRangeQuery(startDate, endDate); dateQuery != "" {
					path += "?" + dateQuery
				}
				resp, err := client.Request("POST", path, nil)
				if err != nil {
					return err
				}
				if resp.Success {
					fmt.Printf("✅ Trainings sync complete: %s\n", resp.Message)
				} else {
					fmt.Printf("❌ Trainings sync failed: %s\n", resp.Message)
					return fmt.Errorf("trainings sync failed")
				}
			}

			if weight {
				resp, err := client.Request(
					"POST", "/api/connect/gymcoach/sync-weight", nil,
				)
				if err != nil {
					return err
				}
				if resp.Success {
					fmt.Printf("✅ %s\n", resp.Message)
				} else {
					fmt.Printf("❌ Weight sync failed: %s\n", resp.Message)
					return fmt.Errorf("weight sync failed")
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&trainings, "trainings", false, "Push GymCoach trainings to Garmin")
	cmd.Flags().BoolVar(&weight, "weight", false, "Push the latest bodyweight to GymCoach")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Backfill from this date (YYYY-MM-DD); requires --end-date")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Backfill through this date (YYYY-MM-DD); requires --start-date")
	return cmd
}
