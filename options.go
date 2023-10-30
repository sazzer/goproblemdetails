package goproblemdetails

// Option is a functional option for customizing the payload of a Problem.
type Option func(*map[string]any)

// WithType is an Option that sets the "type" and "title" properties in the problem payload.
//
// Parameters:
//
//	typeValue - The value for the "type" property.
//	titleValue - The value for the "title" property.
//
// Usage:
//
//	problem := New(http.StatusNotFound, WithType("example-type", "Example Problem Title"))
func WithType(typeValue, titleValue string) Option {
	return func(m *map[string]any) {
		(*m)["type"] = typeValue
		(*m)["title"] = titleValue
	}
}

// WithValue is an Option that sets a custom key-value pair in the problem payload.
//
// Parameters:
//
//	key - The key (property name) for the custom value.
//	value - The custom value associated with the key.
//
// Usage:
//
//	problem := New(http.StatusInternalServerError, WithValue("custom-key", "Custom Value"))
func WithValue(key string, value any) Option {
	return func(m *map[string]any) {
		(*m)[key] = value
	}
}

// WithDetail is an Option that sets the "detail" property in the problem payload.
//
// Parameters:
//
//	value - The value for the "detail" property.
//
// Usage:
//
//	problem := New(http.StatusBadRequest, WithDetail("Additional details about the problem."))
func WithDetail(value string) Option {
	return WithValue("detail", value)
}

// WithInstance is an Option that sets the "instance" property in the problem payload.
//
// Parameters:
//
//	value - The value for the "instance" property.
//
// Usage:
//
//	problem := New(http.StatusConflict, WithInstance("http://example.com/instance/123"))
func WithInstance(value string) Option {
	return WithValue("instance", value)
}
