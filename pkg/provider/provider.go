package provider

type Message interface {
	Role() string
	Content() string
	Usage() (int, int)
}

type Provider interface {
	Name() string
	// SendMessage to a provider
	// TODO: Add tools to this call
	SendMessage(context string, prompt string, messages []Message) Message
}
