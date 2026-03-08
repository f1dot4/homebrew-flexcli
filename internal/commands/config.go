package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/f1dot4/flexcli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func NewConfigCmd(cfgFile *string, rootCfg **config.Config) *cobra.Command {
	var server string
	var key string
	var name string

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure CLI settings for an environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := *rootCfg
			
			if name == "" {
				name = "default"
			}

			ctx := cfg.Contexts[name]
			if server != "" {
				ctx.ServerURL = server
			}

			if key == "" {
				fmt.Printf("Enter FlexCoach API Key for %q: ", name)
				byteKey, err := term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					return err
				}
				fmt.Println()
				key = strings.TrimSpace(string(byteKey))
			}

			if key != "" {
				ctx.APIKey = key
			}

			// Ensure map is initialized
			if cfg.Contexts == nil {
				cfg.Contexts = make(map[string]config.Context)
			}
			cfg.Contexts[name] = ctx
			
			// If first context or specifically configured, set as current
			if cfg.CurrentContext == "" {
				cfg.CurrentContext = name
			}

			if err := config.SaveConfig(*cfgFile, cfg); err != nil {
				return err
			}

			fmt.Printf("Configuration for %q saved securely to %s\n", name, *cfgFile)
			return nil
		},
	}

	cmd.Flags().StringVar(&server, "server", "", "FlexCoach server URL")
	cmd.Flags().StringVar(&key, "key", "", "API Key")
	cmd.Flags().StringVar(&name, "name", "default", "Context name")

	return cmd
}
