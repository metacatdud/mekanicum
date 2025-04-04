package server

type Option func(server *Server)

func WithPromptCapability(listChanged bool) Option {
	return func(s *Server) {
		s.capabilities.prompt = &promptsCapability{
			listChanged: listChanged,
		}
	}
}

func WithResourceCapability(sub, listChanged bool) Option {
	return func(s *Server) {
		s.capabilities.resources = &resourcesCapability{
			listChanged: listChanged,
			subscribed:  sub,
		}
	}
}

func WithToolsCapability(listChanged bool) Option {
	return func(s *Server) {
		s.capabilities.tools = &toolsCapability{
			listChanged: listChanged,
		}
	}
}
