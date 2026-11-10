package errors

import (
	"fmt"
	"time"
	"uuid"

	"github.com/DouglasMaffraValiate/guarana/requestCTX"
)

type Error struct {
	id         uuid.UUID
	instant    time.Time
	code       string
	type_      *ErrorType
	message    string
	context    *ErrorContext
	attributes map[string]string
	cause      error
}

func NewError(ctx requestCTX.Context, errorType *ErrorType, code string, opts ...Option) *Error {
	err := &Error{
		id:      uuid.NewV4(),
		instant: time.Now().UTC(),
		code:    code,
		type_:   errorType,
	}

	err.setContext(ctx)

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(err)
	}
	return err
}

// Options

type Option func(*Error)

func WithAttribute(key string, value any) Option {
	return func(e *Error) {
		e.attributes[key] = fmt.Sprint(value)
	}
}

func WithCause(err error) Option {
	return func(e *Error) {
		e.cause = err
	}
}

// Getters

func (e *Error) GetID() uuid.UUID {
	return e.id
}

func (e *Error) GetInstant() time.Time {
	return e.instant
}

func (e *Error) GetCode() string {
	return e.code
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetAttributes() map[string]string {
	return e.attributes
}

func (e *Error) GetCause() error {
	return e.cause
}
