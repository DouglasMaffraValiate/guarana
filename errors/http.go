package errors

import (
	"net/http"

	"github.com/DouglasMaffraValiate/guarana/requestCTX"
)

var (
	BadRequestError = &ErrorType{
		name:       "BAD_REQUEST",
		httpStatus: http.StatusBadRequest,
	}

	UnauthorizedError = &ErrorType{
		name:       "UNAUTHORIZED",
		httpStatus: http.StatusUnauthorized,
	}

	ForbiddenError = &ErrorType{
		name:       "FORBIDDEN",
		httpStatus: http.StatusForbidden,
	}

	NotFoundError = &ErrorType{
		name:       "NOT_FOUND",
		httpStatus: http.StatusNotFound,
	}

	MethodNotAllowedError = &ErrorType{
		name:       "METHOD_NOT_ALLOWED",
		httpStatus: http.StatusMethodNotAllowed,
	}

	ConflictError = &ErrorType{
		name:       "CONFLICT",
		httpStatus: http.StatusConflict,
	}

	GoneError = &ErrorType{
		name:       "GONE",
		httpStatus: http.StatusGone,
	}

	UnprocessableEntityError = &ErrorType{
		name:       "UNPROCESSABLE_ENTITY",
		httpStatus: http.StatusUnprocessableEntity,
	}

	TooManyRequestsError = &ErrorType{
		name:       "TOO_MANY_REQUESTS",
		httpStatus: http.StatusTooManyRequests,
	}

	InternalServerError = &ErrorType{
		name:       "INTERNAL_SERVER_ERROR",
		httpStatus: http.StatusInternalServerError,
	}

	NotImplementedError = &ErrorType{
		name:       "NOT_IMPLEMENTED",
		httpStatus: http.StatusNotImplemented,
	}

	BadGatewayError = &ErrorType{
		name:       "BAD_GATEWAY",
		httpStatus: http.StatusBadGateway,
	}

	ServiceUnavailableError = &ErrorType{
		name:       "SERVICE_UNAVAILABLE",
		httpStatus: http.StatusServiceUnavailable,
	}

	GatewayTimeoutError = &ErrorType{
		name:       "GATEWAY_TIMEOUT",
		httpStatus: http.StatusGatewayTimeout,
	}
)

func NewBadRequest(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, BadRequestError, code, opts...)
}

func NewUnauthorized(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, UnauthorizedError, code, opts...)
}

func NewForbidden(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, ForbiddenError, code, opts...)
}

func NewNotFound(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, NotFoundError, code, opts...)
}

func NewMethodNotAllowed(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, MethodNotAllowedError, code, opts...)
}

func NewConflict(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, ConflictError, code, opts...)
}

func NewGone(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, GoneError, code, opts...)
}

func NewUnprocessableEntity(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, UnprocessableEntityError, code, opts...)
}

func NewTooManyRequests(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, TooManyRequestsError, code, opts...)
}

func NewInternalServerError(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, InternalServerError, code, opts...)
}

func NewNotImplemented(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, NotImplementedError, code, opts...)
}

func NewBadGateway(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, BadGatewayError, code, opts...)
}

func NewServiceUnavailable(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, ServiceUnavailableError, code, opts...)
}

func NewGatewayTimeout(ctx requestCTX.Context, code string, opts ...Option) *Error {
	return NewError(ctx, GatewayTimeoutError, code, opts...)
}
