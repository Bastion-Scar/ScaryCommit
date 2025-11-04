package cmd

import (
	"ScaryCommit/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

// Creates default config
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize ScaryCommit configuration",
	Long: `Creates a default configuration file for ScaryCommit 
in your home directory (ScaryCommit/deepseekCfg.yaml). 
Use this before your first run to set up provider and API key`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧩 Initializing ScaryCommit configuration...")
		config.SaveDefaultConfig()
		fmt.Println("✅ Configuration initialized successfully!")
		fmt.Println("💡 Edit the file to add your real API key before using scarycommit.")
	},
}
