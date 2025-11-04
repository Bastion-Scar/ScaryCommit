package prompt

import "fmt"

func BuildPrompt(diff, style, lang string) string {
	return fmt.Sprintf("You are an assistant generating a %s commit message in %s.\nDiff:\n%s", style, lang, diff)
}

func TrimDiff(diff string, maxChars int) string {
	if len(diff) > maxChars {
		return diff[:maxChars] + "\n[...diff truncated]"
	}
	return diff
}
