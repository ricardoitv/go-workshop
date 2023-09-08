package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	Client *openai.Client
}

func NewOpenAI(token string) OpenAI {
	return OpenAI{
		Client: openai.NewClient(token),
	}
}

func (o OpenAI) Summarise(transcript string) (string, error) {
	prompt := fmt.Sprintf("Write a synopsis for the following transcript:\n%s", transcript)
	summary, err := o.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", fmt.Errorf("Call to openai failed: %w", err)
	}
	return summary.Choices[0].Message.Content, nil
}
