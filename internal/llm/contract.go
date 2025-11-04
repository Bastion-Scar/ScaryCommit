package llm

import "context"

// GenerateOptions defines tuning parameters for text generation
type GenerateOptions struct {

	// Temperature controls creativity/randomness of output.
	// Example: 0.3 stable and focused; higher → more creativity

	Temperature float64
}

type LLMProvider interface {
	Generate(ctx context.Context, prompt string, opts GenerateOptions) (string, error)
}
