package protocol

type Resource struct {
	URI         string `json:"uri"`                   // Unique identifier (e.g., "file:///path/to/file")
	Name        string `json:"name"`                  // Human-readable name
	Description string `json:"description,omitempty"` // Optional description
	MimeType    string `json:"mimeType,omitempty"`    // Optional MIME type
}

type ResourceOption func(*Resource)

// WithResourceURI sets the URI for the resource.
func WithResourceURI(uri string) ResourceOption {
	return func(r *Resource) {
		r.URI = uri
	}
}

func WithResourceName(name string) ResourceOption {
	return func(r *Resource) {
		r.Name = name
	}
}

func WithResourceDescription(desc string) ResourceOption {
	return func(r *Resource) {
		r.Description = desc
	}
}

func WithResourceMimeType(mime string) ResourceOption {
	return func(r *Resource) {
		r.MimeType = mime
	}
}

func NewResource(opts ...ResourceOption) *Resource {
	resource := &Resource{}
	for _, opt := range opts {
		opt(resource)
	}
	return resource
}
