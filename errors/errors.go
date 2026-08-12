package errors

import "fmt"

// BadRequestError represents a 4xx-style API failure.
type BadRequestError struct {
	Message    string
	StatusCode int
	Body       string
	Err        error
}

func (e *BadRequestError) Error() string {
	return formatError(e.Message, e.StatusCode, e.Err)
}

func (e *BadRequestError) Unwrap() error { return e.Err }

// ServerError represents a 5xx-style API failure.
type ServerError struct {
	Message    string
	StatusCode int
	Body       string
	Err        error
}

func (e *ServerError) Error() string {
	return formatError(e.Message, e.StatusCode, e.Err)
}

func (e *ServerError) Unwrap() error { return e.Err }

// AuthError represents login / token problems.
type AuthError struct {
	Message    string
	StatusCode int
	Body       string
	Err        error
}

func (e *AuthError) Error() string {
	return formatError(e.Message, e.StatusCode, e.Err)
}

func (e *AuthError) Unwrap() error { return e.Err }

func formatError(msg string, status int, err error) string {
	out := "shiprocket: "
	if msg != "" {
		out += msg
	}
	if status > 0 {
		out += fmt.Sprintf(" (status %d)", status)
	}
	if err != nil {
		if msg != "" {
			out += ": "
		}
		out += err.Error()
	}
	return out
}
