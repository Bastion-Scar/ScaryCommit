package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CommitAuto(file, message string) error {
	// Checking changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	out, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if strings.TrimSpace(string(out)) == "" {
		return fmt.Errorf("no changes to commit")
	}

	// Commiting
	var commitOut bytes.Buffer
	commitCmd := exec.Command("git", "commit", file, "-m", message)
	commitCmd.Stdout = &commitOut
	commitCmd.Stderr = &commitOut

	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("failed to create commit: %s (%w)", commitOut.String(), err)
	}

	fmt.Println(commitOut.String())
	return nil
}

func Commit(message string) error {
	// Checking changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	out, err := statusCmd.Output()

	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if strings.TrimSpace(string(out)) == "" {
		return fmt.Errorf("no changes to commit")
	}

	// Commiting
	var commitOut bytes.Buffer
	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Stdout = &commitOut
	commitCmd.Stderr = &commitOut

	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("failed to create commit: %s (%w)", commitOut.String(), err)
	}

	fmt.Println(commitOut.String())
	return nil
}
