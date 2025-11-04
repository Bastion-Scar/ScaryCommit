package git

import (
	"bytes"
	"os/exec"
)

// Executing command git diff --cached --no-color
func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--no-color")
	var out bytes.Buffer // Output buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "Failed to get git diff", err
	}
	return out.String(), nil
}
