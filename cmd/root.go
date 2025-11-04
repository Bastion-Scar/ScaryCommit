package cmd

import (
	"ScaryCommit/internal/ui"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scarycommit",
	Short: "ScaryCommit — AI commit message generator",
	Long:  "Generates commit messages from staged git changes using LLMs like DeepSeek or OpenRouter.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Current commands
func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(commitCmd)
}
