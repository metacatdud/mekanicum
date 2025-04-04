// Package protocol implements the mcp message as described on
// https://modelcontextprotocol.io/introduction
// (this is subject to change)

package protocol

import (
	"encoding/json"
	"fmt"
)

const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)

const (
	JSONRPCVersion = "2.0"
)

// McpMessage will be used for creating both request and response as well as for error
type McpMessage struct {
	JSONRPC   string      `json:"jsonrpc"`
	ID        interface{} `json:"id,omitempty"`
	Method    string      `json:"method,omitempty"`
	Params    interface{} `json:"params,omitempty"`
	Result    interface{} `json:"result,omitempty"`
	Error     *McpError   `json:"error,omitempty"`
	Extension interface{} `json:"extension,omitempty"`
}

func (m *McpMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *McpMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type McpError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRequest(opts ...Option) (*McpMessage, error) {
	msg := NewRawMCPMessage(opts...)

	if msg.ID == nil {
		return nil, fmt.Errorf("request ID cannot be empty")
	}

	if msg.Method == "" {
		return nil, fmt.Errorf("request method cannot be empty")
	}

	return msg, nil
}

func NewResponse(opts ...Option) (*McpMessage, error) {
	msg := NewRawMCPMessage(opts...)

	if msg.ID == nil {
		return nil, fmt.Errorf("request ID cannot be empty")
	}

	if msg.Method == "" {
		return nil, fmt.Errorf("request method cannot be empty")
	}

	return msg, nil
}

func NewErrorResponse(opts ...Option) (*McpMessage, error) {
	msg := NewRawMCPMessage(opts...)
	if msg.ID == nil {
		return nil, fmt.Errorf("request ID cannot be empty")
	}

	if msg.Error == nil {
		return nil, fmt.Errorf("request Error cannot be empty")
	}

	return msg, nil

}

func NewNotification(opts ...Option) (*McpMessage, error) {
	msg := NewRawMCPMessage(opts...)

	if msg.ID == nil {
		return nil, fmt.Errorf("request ID cannot be empty")
	}

	if msg.Method == "" {
		return nil, fmt.Errorf("request method cannot be empty")
	}

	return msg, nil
}

// NewRawMCPMessage creates a message boilerplate
// Should not be used directly but because things might change in the future
// this will allow you to overwrite any of the messages above
func NewRawMCPMessage(opts ...Option) *McpMessage {
	msg := &McpMessage{
		JSONRPC: JSONRPCVersion,
	}

	for _, opt := range opts {
		opt(msg)
	}

	return msg
}
