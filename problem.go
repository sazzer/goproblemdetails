package goproblemdetails

import (
	"encoding/json"
	"fmt"
)

// Problem represents an HTTP Problem as defined by RFC-9457..
// It contains a status code and a payload with additional details.
type Problem struct { //nolint:errname,musttag
	StatusCode int
	Payload    map[string]any
}

// New creates a new instance of the Problem type.
// It allows specifying the HTTP status code and additional options for the problem payload.
//
// Parameters:
//
//	statusCode - The HTTP status code to be associated with the problem.
//	options - A variadic list of Option functions to customize the problem payload.
//
// Returns:
//
//	A Problem instance with the specified status code and payload options.
func New(statusCode int, options ...Option) Problem {
	payload := map[string]any{}

	for _, option := range options {
		option(&payload)
	}

	return Problem{
		StatusCode: statusCode,
		Payload:    payload,
	}
}

func (p Problem) Error() string {
	problemType := p.Payload["type"]

	return fmt.Sprintf("HTTP %d: Problem %v", p.StatusCode, problemType)
}

func (p Problem) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Payload)
}
