package commands

import (
	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewConnectCmd(rootCfg **config.Config, resolvedCtx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Manage device connections and system status",
	}

	cmd.AddCommand(NewStatusCmd(rootCfg, resolvedCtx))
	cmd.AddCommand(NewSyncCmd(rootCfg, resolvedCtx))

	return cmd
}
