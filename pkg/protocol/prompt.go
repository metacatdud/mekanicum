package protocol

type PromptArgument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type Prompt struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Template    string           `json:"template"`
	Arguments   []PromptArgument `json:"arguments,omitempty"`
}

type PromptOption func(*Prompt)

// WithPromptID sets the ID for the prompt.
func WithPromptID(id string) PromptOption {
	return func(p *Prompt) {
		p.ID = id
	}
}

func WithPromptName(name string) PromptOption {
	return func(p *Prompt) {
		p.Name = name
	}
}

func WithPromptDescription(desc string) PromptOption {
	return func(p *Prompt) {
		p.Description = desc
	}
}

func WithPromptTemplate(template string) PromptOption {
	return func(p *Prompt) {
		p.Template = template
	}
}

func WithPromptArguments(args []PromptArgument) PromptOption {
	return func(p *Prompt) {
		p.Arguments = args
	}
}

func NewPrompt(opts ...PromptOption) *Prompt {
	prompt := &Prompt{}
	for _, opt := range opts {
		opt(prompt)
	}
	return prompt
}
