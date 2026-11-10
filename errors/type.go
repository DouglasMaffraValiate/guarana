package errors

import (
	"net/http"
	"strings"
)

// Representa o tipó do Error
type ErrorType struct {
	name       string
	httpStatus uint16
}

// Constructor

// Caso name seja invalido vai subistituir por "INTERNAL_ERROR"
// Caso httpStatus seja invalido vai subistituir por http.StatusInternalServerError
func NewErrorType(name string, httpStatus uint16) *ErrorType {
	er := ErrorType{
		name:       strings.TrimSpace(strings.ToUpper(name)),
		httpStatus: httpStatus,
	}

	if len(er.name) == 0 {
		er.name = "INTERNAL_ERROR"
	}

	if er.httpStatus == 0 || er.httpStatus <= 400 || er.httpStatus >= 500 {
		er.httpStatus = http.StatusInternalServerError
	}

	return &er
}

// Getters

func (e *Error) GetTypeName() string {
	return e.type_.name
}

func (e *Error) GetHTTPStatus() uint16 {
	return e.type_.httpStatus
}
