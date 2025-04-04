package protocol

import "encoding/json"

type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Raw         json.RawMessage        `json:"-"`
}

type ToolOption func(*Tool)

func WithToolName(name string) ToolOption {
	return func(t *Tool) {
		t.Name = name
	}
}

func WithToolDescription(desc string) ToolOption {
	return func(t *Tool) {
		t.Description = desc
	}
}

func WithToolSchema(schema map[string]interface{}) ToolOption {
	return func(t *Tool) {
		t.Schema = schema
	}
}

func NewTool(opts ...ToolOption) *Tool {
	tool := &Tool{}
	for _, opt := range opts {
		opt(tool)
	}
	return tool
}
