package protocol

type Option func(*McpMessage)

func WithID(id interface{}) Option {
	return func(message *McpMessage) {
		message.ID = id
	}
}

func WithMethod(method string) Option {
	return func(message *McpMessage) {
		message.Method = method
	}
}

func WithParams(params map[string]interface{}) Option {
	return func(message *McpMessage) {
		message.Params = params
	}
}

func WithExtension(ext interface{}) Option {
	return func(message *McpMessage) {
		message.Extension = ext
	}
}

func WithResult(result interface{}) Option {
	return func(message *McpMessage) {
		message.Result = result
	}
}

func WithError(code int, msg string, data interface{}) Option {
	return func(message *McpMessage) {
		message.Error = &McpError{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	}
}
