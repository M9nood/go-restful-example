package errors

import (
	"net/http"
	"strconv"
)

var errors = map[string]Errorinfo{
	"ErrorNotFound":       Errorinfo{ErrorCode: "404", ErrorName: "ErrorNotFound"},
	"ErrorBadRequest":     Errorinfo{ErrorCode: "400", ErrorName: "ErrorBadRequest"},
	"ErrorInternalServer": Errorinfo{ErrorCode: "500", ErrorName: "ErrorInternalServer"},
}

type Errorinfo struct {
	ErrorCode    string
	ErrorName    string
	ErrorMessage string `omitempty`
}

type APIError interface {
	Error() string
	GetCode() string
	GetName() string
}

func ErrorHandler(errName string) (code string, name string) {
	if _, ok := errors[errName]; ok {
		return errors[errName].ErrorCode, errors[errName].ErrorName
	}
	return errors["ErrorInternalServer"].ErrorCode, errors["ErrorInternalServer"].ErrorName
}

func NewError(errName string, msg string) *Errorinfo {
	code, name := ErrorHandler(errName)
	return &Errorinfo{
		ErrorCode:    code,
		ErrorName:    name,
		ErrorMessage: msg,
	}
}

func (e *Errorinfo) Error() string {
	if e.ErrorMessage == "" {
		i, _ := strconv.Atoi(e.ErrorCode)
		return http.StatusText(i)
	}
	return e.ErrorMessage
}

func (e *Errorinfo) GetCode() string {
	return e.ErrorCode
}

func (e *Errorinfo) GetName() string {
	return e.ErrorName
}
