package commands

import (
	"fmt"
	"sort"

	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
)

func NewContextCmd(cfgFile *string, rootCfg **config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "Manage environments (contexts)",
	}

	cmd.AddCommand(newContextListCmd(rootCfg))
	cmd.AddCommand(newContextUseCmd(cfgFile, rootCfg))
	cmd.AddCommand(newContextDeleteCmd(cfgFile, rootCfg))

	return cmd
}

func newContextListCmd(rootCfg **config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all contexts",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := *rootCfg
			if len(cfg.Contexts) == 0 {
				fmt.Println("No contexts found. Use 'config' to create one.")
				return
			}

			fmt.Println("CONTEXTS")

			// Sort names for consistent output
			names := make([]string, 0, len(cfg.Contexts))
			for name := range cfg.Contexts {
				names = append(names, name)
			}
			sort.Strings(names)

			for _, name := range names {
				prefix := "  "
				if name == cfg.CurrentContext {
					prefix = "* "
				}
				ctx := cfg.Contexts[name]
				fmt.Printf("%s%s (%s)\n", prefix, name, ctx.ServerURL)
			}
		},
	}
}

func newContextUseCmd(cfgFile *string, rootCfg **config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "use [name]",
		Short: "Switch the active context",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := *rootCfg
			name := args[0]

			if _, ok := cfg.Contexts[name]; !ok {
				return fmt.Errorf("context %q not found", name)
			}

			cfg.CurrentContext = name
			if err := config.SaveConfig(*cfgFile, cfg); err != nil {
				return err
			}

			fmt.Printf("Switched to context %q\n", name)
			return nil
		},
	}
}

func newContextDeleteCmd(cfgFile *string, rootCfg **config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [name]",
		Short: "Remove a context",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := *rootCfg
			name := args[0]

			if _, ok := cfg.Contexts[name]; !ok {
				return fmt.Errorf("context %q not found", name)
			}

			delete(cfg.Contexts, name)
			if cfg.CurrentContext == name {
				cfg.CurrentContext = ""
			}

			if err := config.SaveConfig(*cfgFile, cfg); err != nil {
				return err
			}

			fmt.Printf("Context %q deleted\n", name)
			return nil
		},
	}
}
