package commands

import (
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

// NewProfileDataCmd builds the profile data command tree, grouping
// manual sync, imported activities, imported health metrics, and fitness records.
func NewProfileDataCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "Sync & data: manual sync, activities, health metrics",
	}

	cmd.AddCommand(newDataSyncCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataActivityCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataHealthMetricCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataFitnessCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(newDataNutritionCmd(rootCfg, resolvedCtx))

	return cmd
}
