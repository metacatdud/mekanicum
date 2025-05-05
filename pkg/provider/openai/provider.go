package openai

import (
	"context"
	"errors"
	"mekanicum/pkg/provider"
)

type Provider struct {
	client *Client
	model  string
}

func NewProvider(apiKey, model string) *Provider {
	client := NewClient(apiKey)

	return &Provider{
		client: client,
		model:  model,
	}
}

func (p *Provider) Name() string {
	return "OpenAI"
}

func (p *Provider) SendMessage(ctx context.Context, prompt string, messages []provider.Message) (provider.Message, error) {

	// Prepare the message payload
	messageParams := make([]MessageParam, len(messages))
	for _, msg := range messages {
		messageParam := MessageParam{
			Role: msg.Role(),
		}

		if msg.Content() == "" {
			content := msg.Content()
			messageParam.Content = &content
		}

		messageParams = append(messageParams, messageParam)
	}

	if prompt != "" {
		messageParams = append(messageParams, MessageParam{
			Content: &prompt,
			Role:    "user",
		})
	}

	// TODO: Get and append tools

	// Send request using the client
	responseData, err := p.client.SendRequest(ctx, MessageRequest{
		Model:       p.model,
		Messages:    messageParams,
		Tools:       nil,
		MaxTokens:   4096,
		Temperature: .4,
	})

	if err != nil {
		return nil, err
	}

	// Return the first choice as the response
	if len(responseData.Choices) > 0 {
		return &ProviderMessageResponse{
			ClientResponse: responseData,
			Choice:         &responseData.Choices[0],
		}, nil
	}

	return nil, errors.New("no choices returned in response")
}

type ProviderMessageResponse struct {
	ClientResponse *MessageResponse
	Choice         *Choice
}

func (p *ProviderMessageResponse) Role() string {
	return p.Choice.Message.Role
}

func (p *ProviderMessageResponse) Content() string {
	if p.Choice.Message.Content == nil {
		return ""
	}

	return *p.Choice.Message.Content
}

func (p *ProviderMessageResponse) Usage() (int, int) {
	return p.ClientResponse.Usage.PromptTokens,
		p.ClientResponse.Usage.CompletionTokens
}
