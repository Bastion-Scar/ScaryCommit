package cmd

import (
	"ScaryCommit/internal/config"
	"ScaryCommit/internal/git"
	"ScaryCommit/internal/llm"
	"ScaryCommit/internal/prompt"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Auto-commit all changes using AI (parallel mode)",
	Long: `Analyzes all staged changes, generates commit messages with AI in parallel,
and commits them sequentially to avoid Git concurrency issues.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// Load config
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Add all files to staging
		err = exec.Command("git", "add", "-A").Run()
		if err != nil {
			fmt.Println("❌ Failed to index files")
		}

		diff, err := git.GetDiff()
		if err != nil {
			return fmt.Errorf("failed to get git diff: %w", err)
		}
		if diff == "" {
			fmt.Println("No changes to commit.")
			return nil
		}

		// Split by file
		chunks := prompt.SplitDiffByFile(diff)
		if len(chunks) == 0 {
			fmt.Println("No file changes detected.")
			return nil
		}

		// Init LLM client
		provider := strings.ToLower(cfg.Provider)
		var client llm.LLMProvider
		switch provider {
		case "openrouter":
			client = llm.NewOpenRouter(cfg.APIKey, cfg.Model)
		case "deepseek":
			client = llm.NewDeepSeek(cfg.APIKey, cfg.Model)
		default:
			return fmt.Errorf("unsupported provider: %s", cfg.Provider)
		}

		fmt.Printf("🤖 Generating commit messages for %d files in parallel...\n\n", len(chunks))

		type result struct {
			file    string
			message string
			err     error
		}

		results := make(chan result, len(chunks))
		var wg sync.WaitGroup
		sem := make(chan struct{}, 3)

		// Generating message for every commit
		for file, fileDiff := range chunks {
			wg.Add(1)
			go func(file, fileDiff string) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()

				promptText := prompt.BuildPrompt(fileDiff, cfg.Style, cfg.Language)
				msg, err := client.Generate(ctx, promptText, llm.GenerateOptions{Temperature: 0.3})
				if err != nil {
					results <- result{file, "", fmt.Errorf("AI generation failed: %w", err)}
					return
				}
				results <- result{file, msg, nil}
			}(file, fileDiff)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		// Sorting messages by ❌ and ✅
		var messages []result
		for r := range results {
			if r.err != nil {
				fmt.Printf("❌ [%s] %v\n", r.file, r.err)
			} else {
				fmt.Printf("✅ [%s] message generated\n", r.file)
				messages = append(messages, r)
			}
		}

		// Commiting in order
		for _, r := range messages {
			if !noConfirm {
				fmt.Printf("\nCommit [%s] with message:\n%s\nUse this? (y/N): ", r.file, r.message)
				var confirm string
				_, err := fmt.Scanln(&confirm)
				if err != nil {
					fmt.Printf("❌ Failed to read input: %v\n", err)
					continue
				}

				if confirm != "y" && confirm != "Y" {
					// Убираем файл из индекса
					resetCmd := exec.Command("git", "reset", r.file)
					if err := resetCmd.Run(); err != nil {
						fmt.Printf("⚠️ Failed to unstage [%s]: %v\n", r.file, err)
					} else {
						fmt.Printf("❌ [%s] removed from staging\n", r.file)
					}
					continue
				}
			}

			// Commit confirmed file
			if err := git.Commit(r.file, r.message); err != nil {
				fmt.Printf("❌ [%s] commit failed: %v\n", r.file, err)
			} else {
				fmt.Printf("✅ [%s] committed successfully!\n", r.file)
			}
		}

		fmt.Printf("\n🎯 Done: %d files processed.\n", len(messages))
		return nil
	},
}

func init() {
	autoCmd.Flags().BoolVar(&noConfirm, "no-confirm", false, "Skip confirmation before commiting")
}
