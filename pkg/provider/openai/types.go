package openai

type MessageParam struct {
	Content *string `json:"content"`
	Role    string  `json:"role"`
	Name    string  `json:"name,omitempty"`

	// TODO: Add function call and tool calls here
}

type MessageRequest struct {
	Model       string         `json:"model"`
	Messages    []MessageParam `json:"messages"`
	Tools       []Tool         `json:"tools,omitempty"`
	MaxTokens   int            `json:"max_tokens,omitempty"`
	Temperature float32        `json:"temperature,omitempty"`
}

type MessageResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Tool struct {
	Name string `json:"name"`
}

type Choice struct {
	Message      MessageParam `json:"message"`
	Index        int          `json:"index"`
	FinishReason string       `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
