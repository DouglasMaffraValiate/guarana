package errors

import (
	"errors"
	"fmt"
)

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}

	if e.cause == nil {
		return fmt.Sprintf("[%s] %s", e.code, e.message)
	}

	return fmt.Sprintf("[%s] %s: %v", e.code, e.message, e.cause)
}

// Unwrap permite inspecionar a causa original.
func (e *Error) Unwrap() error {
	if e == nil || e.cause == nil {
		return nil
	}
	return e.cause
}

// Compara erros com base no código e na message
func (e *Error) Is(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}

	return e.code == targetErr.code && e.message == targetErr.message
}

// Extrai um *Error da cadeia
func AsError(err error) (*Error, bool) {
	var target *Error

	if !errors.As(err, &target) {
		return nil, false
	}

	return target, true
}

// Retorna a causa fundamental do erro, ou pelo menos a primeira adicionada
func (e *Error) GetRootCause() error {
	if e == nil {
		return nil
	}

	return e.GetRootCause()
}
