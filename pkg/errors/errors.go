package errors

import "fmt"

type Error struct {
	Code     int                    `json:"code"`
	Reason   string                 `json:"reason"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v", e.Code, e.Reason, e.Message, e.Metadata)
}

// StatusCode return an HTTP error code.
func (e *Error) StatusCode() int {
	return e.Code
}

// Todo
