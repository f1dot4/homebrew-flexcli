package commands

import (
	"encoding/json"
	"fmt"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewGarminCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "garmin",
		Short: "Garmin gamification, equipment, goals, workouts, and devices",
	}

	cmd.AddCommand(newGarminBadgesCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGarminGearCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGarminChallengesCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGarminGoalsCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGarminWorkoutsCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newGarminDevicesCmd(rootCfg, resolvedCtx))

	return cmd
}

func newGarminBadgesCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "badges",
		Short: "List all earned and available Garmin badges",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/badges", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var badges []struct {
				BadgeID       string   `json:"badge_id"`
				BadgeName     *string  `json:"badge_name"`
				BadgeCategory *string  `json:"badge_category"`
				EarnedDate    *string  `json:"earned_date"`
				EarnedNumber  *int     `json:"earned_number"`
				ProgressValue *float64 `json:"progress_value"`
				TargetValue   *float64 `json:"target_value"`
			}
			if err := json.Unmarshal(resp.Data, &badges); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(badges) == 0 {
				fmt.Println("No badges found.")
				return nil
			}

			fmt.Println("🏅 Garmin Badges")
			fmt.Println("================")
			fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-8s\n", "BADGE ID", "NAME", "CATEGORY", "EARNED DATE", "COUNT")
			fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-8s\n", "──────────", "──────────────────────────────", "───────────────", "────────────", "────────")
			for _, b := range badges {
				name := "-"
				if b.BadgeName != nil {
					name = *b.BadgeName
				}
				cat := "-"
				if b.BadgeCategory != nil {
					cat = *b.BadgeCategory
				}
				date := "-"
				if b.EarnedDate != nil {
					date = *b.EarnedDate
				}
				count := "-"
				if b.EarnedNumber != nil {
					count = fmt.Sprintf("%d", *b.EarnedNumber)
				}
				fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-8s\n", b.BadgeID, name, cat, date, count)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGarminGearCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "gear",
		Short: "List all Garmin gear (shoes, bikes, etc.)",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/gear", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var gearList []struct {
				GearUUID       string   `json:"gear_uuid"`
				GearName       *string  `json:"gear_name"`
				GearType       *string  `json:"gear_type"`
				ActivityType   *string  `json:"activity_type"`
				TotalDistanceM *float64 `json:"total_distance_m"`
				TotalDurationS *int     `json:"total_duration_s"`
				IsDefault      bool     `json:"is_default"`
			}
			if err := json.Unmarshal(resp.Data, &gearList); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(gearList) == 0 {
				fmt.Println("No gear found.")
				return nil
			}

			fmt.Println("👟 Garmin Gear")
			fmt.Println("===============")
			fmt.Printf("  %-36s  %-25s  %-15s  %-10s  %-8s\n", "UUID", "NAME", "TYPE", "DISTANCE", "DEFAULT")
			fmt.Printf("  %-36s  %-25s  %-15s  %-10s  %-8s\n", "────────────────────────────────────", "─────────────────────────", "───────────────", "──────────", "────────")
			for _, g := range gearList {
				name := "-"
				if g.GearName != nil {
					name = *g.GearName
				}
				gType := "-"
				if g.GearType != nil {
					gType = *g.GearType
				}
				dist := "0.0 km"
				if g.TotalDistanceM != nil {
					dist = fmt.Sprintf("%.1f km", *g.TotalDistanceM/1000.0)
				}
				def := "No"
				if g.IsDefault {
					def = "Yes"
				}
				fmt.Printf("  %-36s  %-25s  %-15s  %-10s  %-8s\n", g.GearUUID, name, gType, dist, def)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGarminChallengesCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "challenges",
		Short: "List all Garmin challenges",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/challenges", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var challenges []struct {
				ChallengeID   string   `json:"challenge_id"`
				ChallengeName *string  `json:"challenge_name"`
				ChallengeType *string  `json:"challenge_type"`
				Status        *string  `json:"status"`
				StartDate     *string  `json:"start_date"`
				EndDate       *string  `json:"end_date"`
				Progress      *float64 `json:"progress"`
				Target        *float64 `json:"target"`
			}
			if err := json.Unmarshal(resp.Data, &challenges); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(challenges) == 0 {
				fmt.Println("No challenges found.")
				return nil
			}

			fmt.Println("🏆 Garmin Challenges")
			fmt.Println("=====================")
			fmt.Printf("  %-15s  %-30s  %-12s  %-12s  %-15s\n", "CHALLENGE ID", "NAME", "STATUS", "END DATE", "PROGRESS")
			fmt.Printf("  %-15s  %-30s  %-12s  %-12s  %-15s\n", "───────────────", "──────────────────────────────", "────────────", "────────────", "───────────────")
			for _, c := range challenges {
				name := "-"
				if c.ChallengeName != nil {
					name = *c.ChallengeName
				}
				status := "-"
				if c.Status != nil {
					status = *c.Status
				}
				endDate := "-"
				if c.EndDate != nil {
					endDate = *c.EndDate
				}
				prog := "-"
				if c.Progress != nil && c.Target != nil {
					prog = fmt.Sprintf("%.1f / %.1f", *c.Progress, *c.Target)
				}
				fmt.Printf("  %-15s  %-30s  %-12s  %-12s  %-15s\n", c.ChallengeID, name, status, endDate, prog)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGarminGoalsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "goals",
		Short: "List all Garmin training goals",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/goals", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var goals []struct {
				GoalID       string   `json:"goal_id"`
				GoalName     *string  `json:"goal_name"`
				GoalType     *string  `json:"goal_type"`
				StartDate    *string  `json:"start_date"`
				TargetValue  *float64 `json:"target_value"`
				CurrentValue *float64 `json:"current_value"`
				Unit         *string  `json:"unit"`
			}
			if err := json.Unmarshal(resp.Data, &goals); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(goals) == 0 {
				fmt.Println("No Garmin goals found.")
				return nil
			}

			fmt.Println("🎯 Garmin Goals")
			fmt.Println("================")
			fmt.Printf("  %-10s  %-25s  %-15s  %-12s  %-15s\n", "GOAL ID", "NAME", "TYPE", "START DATE", "PROGRESS")
			fmt.Printf("  %-10s  %-25s  %-15s  %-12s  %-15s\n", "──────────", "─────────────────────────", "───────────────", "────────────", "───────────────")
			for _, g := range goals {
				name := "-"
				if g.GoalName != nil {
					name = *g.GoalName
				}
				gType := "-"
				if g.GoalType != nil {
					gType = *g.GoalType
				}
				startDate := "-"
				if g.StartDate != nil {
					startDate = *g.StartDate
				}
				prog := "-"
				if g.CurrentValue != nil && g.TargetValue != nil {
					unit := ""
					if g.Unit != nil {
						unit = " " + *g.Unit
					}
					prog = fmt.Sprintf("%.1f / %.1f%s", *g.CurrentValue, *g.TargetValue, unit)
				}
				fmt.Printf("  %-10s  %-25s  %-15s  %-12s  %-15s\n", g.GoalID, name, gType, startDate, prog)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGarminWorkoutsCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "workouts",
		Short: "List all Garmin scheduled workouts",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/workouts", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var workouts []struct {
				WorkoutID          string   `json:"workout_id"`
				WorkoutName        *string  `json:"workout_name"`
				SportType          *string  `json:"sport_type"`
				EstimatedDurationS *int     `json:"estimated_duration_s"`
				EstimatedDistanceM *float64 `json:"estimated_distance_m"`
			}
			if err := json.Unmarshal(resp.Data, &workouts); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(workouts) == 0 {
				fmt.Println("No workouts found.")
				return nil
			}

			fmt.Println("📋 Garmin Workouts")
			fmt.Println("===================")
			fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-10s\n", "ID", "NAME", "SPORT", "DURATION", "DISTANCE")
			fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-10s\n", "──────────", "──────────────────────────────", "───────────────", "────────────", "──────────")
			for _, w := range workouts {
				name := "-"
				if w.WorkoutName != nil {
					name = *w.WorkoutName
				}
				sport := "-"
				if w.SportType != nil {
					sport = *w.SportType
				}
				dur := "-"
				if w.EstimatedDurationS != nil {
					m := *w.EstimatedDurationS / 60
					dur = fmt.Sprintf("%d min", m)
				}
				dist := "-"
				if w.EstimatedDistanceM != nil {
					dist = fmt.Sprintf("%.1f km", *w.EstimatedDistanceM/1000.0)
				}
				fmt.Printf("  %-10s  %-30s  %-15s  %-12s  %-10s\n", w.WorkoutID, name, sport, dur, dist)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}

func newGarminDevicesCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "devices",
		Short: "List all registered Garmin devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("GET", "/api/garmin/devices", nil)
			if err != nil {
				return err
			}
			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			var devices []struct {
				DeviceID        string  `json:"device_id"`
				DeviceName      *string `json:"device_name"`
				DeviceType      *string `json:"device_type"`
				FirmwareVersion *string `json:"firmware_version"`
				BatteryLevel    *int    `json:"battery_level"`
			}
			if err := json.Unmarshal(resp.Data, &devices); err != nil {
				return fmt.Errorf("failed to parse response: %w", err)
			}
			if len(devices) == 0 {
				fmt.Println("No devices found.")
				return nil
			}

			fmt.Println("📱 Garmin Devices")
			fmt.Println("==================")
			fmt.Printf("  %-15s  %-25s  %-15s  %-10s  %-8s\n", "DEVICE ID", "NAME", "TYPE", "FIRMWARE", "BATTERY")
			fmt.Printf("  %-15s  %-25s  %-15s  %-10s  %-8s\n", "───────────────", "─────────────────────────", "───────────────", "──────────", "────────")
			for _, d := range devices {
				name := "-"
				if d.DeviceName != nil {
					name = *d.DeviceName
				}
				dType := "-"
				if d.DeviceType != nil {
					dType = *d.DeviceType
				}
				fw := "-"
				if d.FirmwareVersion != nil {
					fw = *d.FirmwareVersion
				}
				bat := "-"
				if d.BatteryLevel != nil {
					bat = fmt.Sprintf("%d%%", *d.BatteryLevel)
				}
				fmt.Printf("  %-15s  %-25s  %-15s  %-10s  %-8s\n", d.DeviceID, name, dType, fw, bat)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")
	return cmd
}
