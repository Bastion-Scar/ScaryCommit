package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DeepSeekClient struct {
	APIKey string
	Model  string
	URL    string
}

// Constructor [lego :)]
func NewDeepSeek(apiKey, model string) *DeepSeekClient {
	return &DeepSeekClient{
		APIKey: apiKey,
		Model:  model,
		URL:    "https://api.deepseek.com/chat/completions",
	}
}

// Chat API structures

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model       string        `json:"model"`
	Messages    []chatMessage `json:"messages"`
	Stream      bool          `json:"stream"`
	Temperature float64       `json:"temperature"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Sending request to DeepSeek API

func (d *DeepSeekClient) Generate(ctx context.Context, prompt string, opts GenerateOptions) (string, error) {
	if opts.Temperature == 0 {
		opts.Temperature = 0.3
	}
	// Prompt
	reqData := chatRequest{
		Model: d.Model,
		Messages: []chatMessage{
			{Role: "system", Content: "You are to act as an author of a commit message in git. DONT use GitMoji or any emoji. Commits MUST be SHORT"},
			{Role: "user", Content: prompt},
		},
		Stream:      false,
		Temperature: opts.Temperature,
	}

	body, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	// Sending -d prompt to API
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.URL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	// Adding -H Headers for request to API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.APIKey)

	client := &http.Client{Timeout: 15 * time.Second} // Timeout 15sec
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(data))
	}

	// Getting response from API
	var out chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if len(out.Choices) == 0 || out.Choices[0].Message.Content == "" {
		return "", errors.New("empty response from DeepSeek")
	}

	return out.Choices[0].Message.Content, nil
}
