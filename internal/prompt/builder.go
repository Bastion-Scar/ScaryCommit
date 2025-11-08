package prompt

import (
	"fmt"
	"strings"
)

func BuildPrompt(diff, style, lang string) string {
	return fmt.Sprintf("You are an assistant generating a %s commit message in %s.\nDiff:\n%s", style, lang, diff)
}

func TrimDiff(diff string, maxChars int) string {
	if len(diff) > maxChars {
		return diff[:maxChars] + "\n[...diff truncated]"
	}
	return diff
}

func SplitDiffByFile(diff string) map[string]string {
	chunks := make(map[string]string)
	lines := strings.Split(diff, "\n")
	var currentFile string
	var currentChunk strings.Builder

	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			if currentFile != "" {
				chunks[currentFile] = currentChunk.String()
				currentChunk.Reset()
			}
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				currentFile = strings.TrimPrefix(parts[2], "b/")
				currentFile = strings.TrimPrefix(currentFile, "a/")
			}
		} else {
			currentChunk.WriteString(line + "\n")
		}
	}

	if currentFile != "" {
		chunks[currentFile] = currentChunk.String()
	}

	return chunks
}
