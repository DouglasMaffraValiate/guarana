package errors

import "github.com/DouglasMaffraValiate/guarana/requestCTX"

// Representa os atributos do requestCTX.Context que fazem sentido para um erro
// Valores imutaveis uma vez instânciados
type ErrorContext struct {
	requestID   string
	tenantID    string
	principalID string

	httpMethod string
	httpPath   string
}

// Getters

func (e *Error) GetRequestID() string {
	return e.context.requestID
}

func (e *Error) GetTenantID() string {
	return e.context.tenantID
}

func (e *Error) GetPrincipalID() string {
	return e.context.principalID
}

func (e *Error) GetHTTPMethod() string {
	return e.context.httpMethod
}

func (e *Error) GetHTTPPath() string {
	return e.context.httpPath
}

// Setter

func (e *Error) setContext(ctx requestCTX.Context) {
	mt := ctx.Metadata()

	e.context = &ErrorContext{
		requestID:  mt.RequestID,
		tenantID:   mt.TenantID,
		httpMethod: ctx.HTTPMethod(),
		httpPath:   ctx.HTTPPath(),
	}

	pr, ok := ctx.Principal()

	if ok {
		e.context.principalID = pr.PrincipalID
	}

}
