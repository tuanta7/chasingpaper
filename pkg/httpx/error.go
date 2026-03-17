package httpx

import "net/http"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Hint    string `json:"hint,omitempty"`
}

func NewError(code int, message string, opts ...Option) Error {
	err := Error{
		Code:    code,
		Message: message,
	}

	for _, opt := range opts {
		opt(&err)
	}

	return err
}

type Option func(e *Error)

func WithHint(hint string) Option {
	return func(e *Error) {
		e.Hint = hint
	}
}

func WithMessage(message string) Option {
	return func(e *Error) {
		e.Message = message
	}
}

func (he Error) Error() string {
	return he.Message
}

func NewInvalidArgumentError(opts ...Option) Error {
	return NewError(http.StatusBadRequest, "invalid argument", opts...)
}

func NewUnauthorizedError(opts ...Option) Error {
	return NewError(http.StatusUnauthorized, "unauthorized", opts...)
}

func NewForbiddenError(opts ...Option) Error {
	return NewError(http.StatusForbidden, "forbidden", opts...)
}

func NewNotFoundError(opts ...Option) Error {
	return NewError(http.StatusNotFound, "not found", opts...)
}

func NewInternalError(opts ...Option) Error {
	return NewError(http.StatusInternalServerError, "internal error", opts...)
}
