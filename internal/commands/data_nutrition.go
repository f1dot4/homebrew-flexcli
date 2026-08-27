package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/f1dot4/flexcli/internal/api"
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// data nutrition (alias: nut)
// ---------------------------------------------------------------------------

func newDataNutritionCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nutrition",
		Aliases: []string{"nut"},
		Short:   "Manage nutrition logs (alias: nut): log, list",
	}

	cmd.AddCommand(newDataNutritionLogCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataNutritionListCmd(rootCfg, resolvedCtx))

	return cmd
}

func newDataNutritionLogCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var (
		calories int
		carbs    float64
		protein  float64
		fat      float64
		sugar    float64
		sodium   int
		fiber    float64
		name     string
		meal     string
		notes    string
		timeStr  string
	)

	cmd := &cobra.Command{
		Use:   "log",
		Short: "Log a food/nutrition entry",
		RunE: func(cmd *cobra.Command, args []string) error {
			meal = strings.ToLower(strings.TrimSpace(meal))
			if meal != "" {
				validMeals := map[string]bool{
					"breakfast": true,
					"lunch":     true,
					"dinner":    true,
					"snack":     true,
				}
				if !validMeals[meal] {
					return fmt.Errorf("invalid meal %q: must be one of breakfast, lunch, dinner, snack", meal)
				}
			}

			loggedAt, err := parseNutritionTime(timeStr)
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"logged_at": loggedAt,
			}
			if cmd.Flags().Changed("calories") {
				body["calories"] = calories
			}
			if cmd.Flags().Changed("carbs") {
				body["carbs_g"] = carbs
			}
			if cmd.Flags().Changed("protein") {
				body["protein_g"] = protein
			}
			if cmd.Flags().Changed("fat") {
				body["fat_g"] = fat
			}
			if cmd.Flags().Changed("sugar") {
				body["sugar_g"] = sugar
			}
			if cmd.Flags().Changed("sodium") {
				body["sodium_mg"] = sodium
			}
			if cmd.Flags().Changed("fiber") {
				body["fiber_g"] = fiber
			}
			if name != "" {
				body["name"] = name
			}
			if meal != "" {
				body["meal_type"] = meal
			}
			if notes != "" {
				body["notes"] = notes
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)
			resp, err := client.Request("POST", "/api/nutrition-log", body)
			if err != nil {
				return err
			}

			if resp.Success {
				fmt.Println("✅ Nutrition entry logged successfully.")
			} else {
				fmt.Printf("❌ Failed to log nutrition entry: %s\n", resp.Message)
			}
			return nil
		},
	}

	cmd.Flags().IntVar(&calories, "calories", 0, "Calories (kcal)")
	cmd.Flags().Float64Var(&carbs, "carbs", 0, "Carbohydrates (g)")
	cmd.Flags().Float64Var(&protein, "protein", 0, "Protein (g)")
	cmd.Flags().Float64Var(&fat, "fat", 0, "Fat (g)")
	cmd.Flags().Float64Var(&sugar, "sugar", 0, "Sugar (g)")
	cmd.Flags().IntVar(&sodium, "sodium", 0, "Sodium (mg)")
	cmd.Flags().Float64Var(&fiber, "fiber", 0, "Fiber (g)")
	cmd.Flags().StringVar(&name, "name", "", "Name of the food item or meal")
	cmd.Flags().StringVar(&meal, "meal", "", "Meal type (breakfast, lunch, dinner, snack)")
	cmd.Flags().StringVar(&notes, "notes", "", "Optional notes")
	cmd.Flags().StringVar(&timeStr, "time", "", "Time of entry (YYYY-MM-DD or YYYY-MM-DD HH:MM, default: now)")

	return cmd
}

func parseNutritionTime(timeStr string) (string, error) {
	if timeStr == "" {
		return time.Now().Format("2006-01-02T15:04:05"), nil
	}
	if t, err := time.Parse("2006-01-02 15:04", timeStr); err == nil {
		return t.Format("2006-01-02T15:04:05"), nil
	}
	if t, err := time.Parse("2006-01-02T15:04", timeStr); err == nil {
		return t.Format("2006-01-02T15:04:05"), nil
	}
	if t, err := time.Parse("2006-01-02", timeStr); err == nil {
		return t.Format("2006-01-02T15:04:05"), nil
	}
	if t, err := time.Parse("2006-01-02T15:04:05", timeStr); err == nil {
		return t.Format("2006-01-02T15:04:05"), nil
	}
	return "", fmt.Errorf("invalid --time format %q: expected YYYY-MM-DD or YYYY-MM-DD HH:MM", timeStr)
}

func newDataNutritionListCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	var (
		startDate string
		endDate   string
		days      int
		asJSON    bool
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List nutrition entries and daily totals",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("days") && (startDate != "" || endDate != "") {
				return fmt.Errorf("cannot specify both --days and --start-date/--end-date")
			}
			if err := validateDateFlag(startDate, "--start-date"); err != nil {
				return err
			}
			if err := validateDateFlag(endDate, "--end-date"); err != nil {
				return err
			}

			client := api.NewClient(resolvedCtx.ServerURL, resolvedCtx.APIKey)

			path := "/api/nutrition-log"
			if cmd.Flags().Changed("days") {
				path += fmt.Sprintf("?days=%d", days)
			} else if q := buildDateRangeQuery(startDate, endDate); q != "" {
				path += "?" + q
			}

			resp, err := client.Request("GET", path, nil)
			if err != nil {
				return err
			}

			if asJSON {
				fmt.Println(string(resp.Data))
				return nil
			}

			return renderNutritionList(resp.Data)
		},
	}

	cmd.Flags().StringVar(&startDate, "start-date", "", "Only include entries on or after this date (YYYY-MM-DD)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Only include entries on or before this date (YYYY-MM-DD)")
	cmd.Flags().IntVarP(&days, "days", "d", 0, "Number of days to list up to today")
	cmd.Flags().BoolVar(&asJSON, "json", false, "Output as JSON")

	return cmd
}

type nutritionEntryDTO struct {
	ID       string   `json:"id"`
	UserID   string   `json:"user_id"`
	LoggedAt string   `json:"logged_at"`
	Name     *string  `json:"name"`
	MealType *string  `json:"meal_type"`
	Calories *int     `json:"calories"`
	CarbsG   *float64 `json:"carbs_g"`
	ProteinG *float64 `json:"protein_g"`
	FatG     *float64 `json:"fat_g"`
	SugarG   *float64 `json:"sugar_g"`
	SodiumMg *int     `json:"sodium_mg"`
	FiberG   *float64 `json:"fiber_g"`
	Notes    *string  `json:"notes"`
}

type nutritionDailyTotalDTO struct {
	Date     string  `json:"date"`
	Calories int     `json:"calories"`
	CarbsG   float64 `json:"carbs_g"`
	ProteinG float64 `json:"protein_g"`
	FatG     float64 `json:"fat_g"`
	SugarG   float64 `json:"sugar_g"`
	SodiumMg int     `json:"sodium_mg"`
	FiberG   float64 `json:"fiber_g"`
}

func renderNutritionList(dataBytes []byte) error {
	var data struct {
		Entries     []nutritionEntryDTO      `json:"entries"`
		DailyTotals []nutritionDailyTotalDTO `json:"daily_totals"`
	}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	entriesByDate := make(map[string][]nutritionEntryDTO)
	for _, entry := range data.Entries {
		dayStr := entry.LoggedAt
		if len(dayStr) >= 10 {
			dayStr = dayStr[:10]
		}
		entriesByDate[dayStr] = append(entriesByDate[dayStr], entry)
	}

	for i, total := range data.DailyTotals {
		if i > 0 {
			fmt.Println()
		}
		dayEntries := entriesByDate[total.Date]
		renderDayBlock(total.Date, dayEntries, total)
	}

	return nil
}

func renderDayBlock(day string, entries []nutritionEntryDTO, total nutritionDailyTotalDTO) {
	fmt.Printf("Nutrition log (%s, %d entries):\n\n", day, len(entries))
	if len(entries) == 0 {
		fmt.Println("No entries.")
		return
	}

	fmt.Printf("  %-5s  %-17s  %-9s  %4s  %6s  %5s  %5s  %5s  %6s  %5s\n",
		"TIME", "NAME", "MEAL", "CAL", "CARB", "PROT", "FAT", "SUGAR", "SODIUM", "FIBER")
	fmt.Printf("  %-5s  %-17s  %-9s  %4s  %6s  %5s  %5s  %5s  %6s  %5s\n",
		"─────", "─────────────────", "─────────", "───", "─────", "─────", "─────", "─────", "──────", "─────")

	for _, e := range entries {
		timeStr := ""
		if len(e.LoggedAt) >= 16 {
			timeStr = e.LoggedAt[11:16]
		}
		nameStr := "-"
		if e.Name != nil && *e.Name != "" {
			nameStr = *e.Name
		}
		mealStr := "-"
		if e.MealType != nil && *e.MealType != "" {
			mealStr = *e.MealType
		}
		calStr := "-"
		if e.Calories != nil {
			calStr = fmt.Sprintf("%d", *e.Calories)
		}
		carbStr := "-"
		if e.CarbsG != nil {
			carbStr = fmt.Sprintf("%5.1fg", *e.CarbsG)
		}
		protStr := "-"
		if e.ProteinG != nil {
			protStr = fmt.Sprintf("%5.1fg", *e.ProteinG)
		}
		fatStr := "-"
		if e.FatG != nil {
			fatStr = fmt.Sprintf("%5.1fg", *e.FatG)
		}
		sugStr := "-"
		if e.SugarG != nil {
			sugStr = fmt.Sprintf("%5.1fg", *e.SugarG)
		}
		sodStr := "-"
		if e.SodiumMg != nil {
			sodStr = fmt.Sprintf("%5dmg", *e.SodiumMg)
		}
		fibStr := "-"
		if e.FiberG != nil {
			fibStr = fmt.Sprintf("%5.1fg", *e.FiberG)
		}

		fmt.Printf("  %-5s  %-17s  %-9s  %4s  %6s  %5s  %5s  %5s  %6s  %5s\n",
			timeStr, nameStr, mealStr, calStr, carbStr, protStr, fatStr, sugStr, sodStr, fibStr)
	}

	fmt.Printf("  %-5s  %-17s  %-9s  %4s  %6s  %5s  %5s  %5s  %6s  %5s\n",
		"─────", "─────────────────", "─────────", "───", "─────", "─────", "─────", "─────", "──────", "─────")
	fmt.Printf("  %-5s  %-17s  %-9s  %4d  %5.1fg  %5.1fg  %5.1fg  %5.1fg  %5dmg  %5.1fg\n",
		"", "", "TOTAL", total.Calories, total.CarbsG, total.ProteinG, total.FatG, total.SugarG, total.SodiumMg, total.FiberG)
}
