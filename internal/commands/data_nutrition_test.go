package commands

import (
	"strings"
	"testing"

	"github.com/f1dot4/flexcli/internal/config"
)

func TestParseNutritionTime(t *testing.T) {
	// Empty string defaults to current timestamp
	nowStr, err := parseNutritionTime("")
	if err != nil {
		t.Fatalf("expected no error for empty time, got: %v", err)
	}
	if len(nowStr) != 19 {
		t.Errorf("expected 19-char ISO timestamp, got: %q", nowStr)
	}

	// YYYY-MM-DD
	dStr, err := parseNutritionTime("2026-08-27")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dStr != "2026-08-27T00:00:00" {
		t.Errorf("expected 2026-08-27T00:00:00, got %q", dStr)
	}

	// YYYY-MM-DD HH:MM
	dtStr, err := parseNutritionTime("2026-08-27 13:02")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dtStr != "2026-08-27T13:02:00" {
		t.Errorf("expected 2026-08-27T13:02:00, got %q", dtStr)
	}

	// Invalid format
	_, err = parseNutritionTime("invalid-time")
	if err == nil {
		t.Fatalf("expected error for invalid time format, got nil")
	}
}

func TestNewDataNutritionCmdTree(t *testing.T) {
	var cfg *config.Config
	ctx := &config.Context{
		ServerURL: "http://localhost:8000",
		APIKey:    "test-key",
	}

	cmd := newDataNutritionCmd(&cfg, ctx)
	if cmd.Use != "nutrition" {
		t.Errorf("expected Use 'nutrition', got %q", cmd.Use)
	}
	if len(cmd.Aliases) != 1 || cmd.Aliases[0] != "nut" {
		t.Errorf("expected alias 'nut', got %v", cmd.Aliases)
	}

	subCommands := cmd.Commands()
	if len(subCommands) != 4 {
		t.Fatalf("expected 4 subcommands, got %d", len(subCommands))
	}

	names := map[string]bool{}
	for _, sub := range subCommands {
		names[sub.Use] = true
	}
	if !names["log"] || !names["list"] || !names["delete <id>"] || !names["restore <id>"] {
		t.Errorf("expected subcommands 'log', 'list', 'delete <id>', and 'restore <id>', got %v", names)
	}
}

func TestNewDataNutritionRestoreCmd(t *testing.T) {
	var cfg *config.Config
	ctx := &config.Context{
		ServerURL: "http://localhost:8000",
		APIKey:    "test-key",
	}

	restoreCmd := newDataNutritionRestoreCmd(&cfg, ctx)
	if restoreCmd.Use != "restore <id>" {
		t.Errorf("expected Use 'restore <id>', got %q", restoreCmd.Use)
	}
	if !restoreCmd.Hidden {
		t.Errorf("expected restore command to be Hidden")
	}

	// Verify exact args = 1
	if err := restoreCmd.Args(restoreCmd, []string{}); err == nil {
		t.Errorf("expected error for 0 args, got nil")
	}
	if err := restoreCmd.Args(restoreCmd, []string{"id1", "id2"}); err == nil {
		t.Errorf("expected error for 2 args, got nil")
	}
	if err := restoreCmd.Args(restoreCmd, []string{"id1"}); err != nil {
		t.Errorf("expected nil error for 1 arg, got %v", err)
	}
}

func TestDataNutritionCmd_RestoreHiddenFromHelp(t *testing.T) {
	var cfg *config.Config
	ctx := &config.Context{
		ServerURL: "http://localhost:8000",
		APIKey:    "test-key",
	}

	cmd := newDataNutritionCmd(&cfg, ctx)
	usage := cmd.UsageString()
	if strings.Contains(usage, "restore") {
		t.Errorf("expected 'restore' to be absent from usage/help text, got: %s", usage)
	}
}

func TestNewDataNutritionDeleteCmd(t *testing.T) {
	var cfg *config.Config
	ctx := &config.Context{
		ServerURL: "http://localhost:8000",
		APIKey:    "test-key",
	}

	delCmd := newDataNutritionDeleteCmd(&cfg, ctx)
	if delCmd.Use != "delete <id>" {
		t.Errorf("expected Use 'delete <id>', got %q", delCmd.Use)
	}

	flag := delCmd.Flags().Lookup("hard")
	if flag == nil {
		t.Fatalf("expected flag 'hard' to exist")
	}
	if !flag.Hidden {
		t.Errorf("expected flag 'hard' to be hidden")
	}

	// Verify exact args = 1
	if err := delCmd.Args(delCmd, []string{}); err == nil {
		t.Errorf("expected error for 0 args, got nil")
	}
	if err := delCmd.Args(delCmd, []string{"id1", "id2"}); err == nil {
		t.Errorf("expected error for 2 args, got nil")
	}
	if err := delCmd.Args(delCmd, []string{"id1"}); err != nil {
		t.Errorf("expected nil error for 1 arg, got %v", err)
	}
}

func TestNewDataNutritionLogCmd_EmptyFlagsRejected(t *testing.T) {
	var cfg *config.Config
	ctx := &config.Context{
		ServerURL: "http://localhost:8000",
		APIKey:    "test-key",
	}

	logCmd := newDataNutritionLogCmd(&cfg, ctx)
	// Running with no flags should return an error
	err := logCmd.RunE(logCmd, []string{})
	if err == nil {
		t.Fatalf("expected error when running nutrition log with no content flags, got nil")
	}
	if !strings.Contains(err.Error(), "at least one parameter must be provided") {
		t.Errorf("expected error message to mention at least one parameter, got: %v", err)
	}
}

func TestRenderNutritionList(t *testing.T) {
	jsonData := []byte(`{
		"entries": [
			{
				"id": "e1",
				"user_id": "u1",
				"logged_at": "2026-08-27T08:15:00",
				"name": "Oatmeal + berries",
				"meal_type": "breakfast",
				"calories": 420,
				"carbs_g": 65.0,
				"protein_g": 12.0,
				"fat_g": 8.0,
				"sugar_g": 18.0,
				"sodium_mg": 150,
				"fiber_g": 9.0
			}
		],
		"daily_totals": [
			{
				"date": "2026-08-27",
				"calories": 420,
				"carbs_g": 65.0,
				"protein_g": 12.0,
				"fat_g": 8.0,
				"sugar_g": 18.0,
				"sodium_mg": 150,
				"fiber_g": 9.0
			},
			{
				"date": "2026-08-26",
				"calories": 0,
				"carbs_g": 0.0,
				"protein_g": 0.0,
				"fat_g": 0.0,
				"sugar_g": 0.0,
				"sodium_mg": 0,
				"fiber_g": 0.0
			}
		]
	}`)

	err := renderNutritionList(jsonData)
	if err != nil {
		t.Fatalf("expected no error rendering nutrition list, got: %v", err)
	}

	err = renderNutritionList([]byte(`invalid json`))
	if err == nil {
		t.Fatalf("expected error for invalid json, got nil")
	}
}

func TestRenderNutritionList_StringNumbers(t *testing.T) {
	// Pydantic Decimal serialization produces strings in JSON mode
	jsonData := []byte(`{
		"entries": [
			{
				"id": "e1",
				"user_id": "u1",
				"logged_at": "2026-08-27T08:15:00",
				"name": "Oatmeal",
				"meal_type": "breakfast",
				"calories": 420,
				"carbs_g": "65.0",
				"protein_g": "12.0",
				"fat_g": "8.0",
				"sugar_g": "18.0",
				"sodium_mg": 150,
				"fiber_g": "9.0"
			}
		],
		"daily_totals": [
			{
				"date": "2026-08-27",
				"calories": 420,
				"carbs_g": "65.0",
				"protein_g": "12.0",
				"fat_g": "8.0",
				"sugar_g": "18.0",
				"sodium_mg": 150,
				"fiber_g": "9.0"
			}
		]
	}`)

	err := renderNutritionList(jsonData)
	if err != nil {
		t.Fatalf("expected no error rendering nutrition list with string floats, got: %v", err)
	}
}
