package main

import (
	"context"
	"fmt"
	"mekanicum/pkg/provider"
	"mekanicum/pkg/provider/openai"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	model := "gpt-4o"
	llm := openai.NewProvider(apiKey, model)

	prompt := "Hello, how are you?"
	messages := make([]provider.Message, 0)

	response, err := llm.SendMessage(context.Background(), prompt, messages)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Response:", response.Content())
}
