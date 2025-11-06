package cmd

import (
	"ScaryCommit/internal/config"
	"ScaryCommit/internal/git"
	"ScaryCommit/internal/llm"
	"ScaryCommit/internal/prompt"
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate and create an AI commit message",
	Long: `Analyzes your staged git changes, sends them to an AI model 
(DeepSeek, OpenRouter, etc.), and creates a commit with the generated message.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// Loading cfg
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		if cfg.APIKey == "" {
			return fmt.Errorf("API key not found — please edit your config file and add it")
		}

		// Getting gitdiff
		diff, err := git.GetDiff()
		if err != nil {
			return fmt.Errorf("failed to get git diff: %w", err)
		}
		if diff == "" {
			return fmt.Errorf("no staged changes — use 'git add' first")
		}

		// User prompt
		trimmedDiff := prompt.TrimDiff(diff, 4000)
		promptToAI := prompt.BuildPrompt(trimmedDiff, cfg.Style, cfg.Language)

		// Initing llm (will be MORE)
		var client llm.LLMProvider
		if cfg.Provider == strings.ToLower("openrouter") {
			client = llm.NewOpenRouter(cfg.APIKey, cfg.Model)
		}
		if cfg.Provider == strings.ToLower("deepseek") {
			client = llm.NewDeepSeek(cfg.APIKey, cfg.Model)
		} else {
			fmt.Println("Incorrect input or provider is not yet supported")
		}

		// Generating msg
		message, err := client.Generate(ctx, promptToAI, llm.GenerateOptions{
			Temperature: 0.3,
		})
		if err != nil {
			return fmt.Errorf("AI generation failed: %w", err)
		}

		fmt.Println("🤖 Suggested commit message:\n")
		fmt.Println(message)
		fmt.Println()

		// Asking user for confirm
		fmt.Print("Do you want to use this message? (y/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			fmt.Println("❌ Commit canceled.")
			return nil
		}

		// Creating commit
		if err := git.Commit(message); err != nil {
			return fmt.Errorf("git commit failed: %w", err)
		}

		fmt.Println("✅ Commit created successfully!")
		return nil
	},
}
