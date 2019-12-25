package errors

import "fmt"

// Error contains errors message and code
type Error struct {
	prefix string
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"message,omitempty"`
}

// New returns new Error
func New(msg string) *Error {
	return &Error{Msg: msg}
}

// WithCode returns Error with errors code
func (e *Error) WithCode(code int) *Error {
	e.Code = code
	return e
}

// WithPrefix returns Error with prefix string
func (e *Error) WithPrefix(prefix string) *Error {
	e.prefix = prefix
	return e
}

// Is returns true when Errors code equal arguments code
func (e *Error) Is(code int) bool {
	return e.Code == code
}

// Error returns Error as string
func (e *Error) Error() string {
	var prefix string
	if e.prefix != "" {
		prefix = e.prefix + ": "
	}
	if e.Code != 0 {
		return fmt.Sprintf("%s%s (code: %d)", prefix, e.Msg, e.Code)
	}
	return e.Msg
}
