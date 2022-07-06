package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Code     int         `json:"code"`
	E        string      `json:"e"`
	Overview interface{} `json:"overview"`
}

func NewError(code int, e error, Overview string) []Error {
	var ve validator.ValidationErrors
	if errors.As(e, &ve) {
		errs := make([]Error, len(ve))
		for i, fe := range ve {
			errs[i] = Error{Code: code, E: e.Error(), Overview: map[string]interface{}{"field": fe.Field(), "message": getBindingErrorMsg(fe)}}
		}
		return errs
	}
	return []Error{{code, e.Error(), Overview}}
}

func getBindingErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required!"
	case "lte":
		return "This field should be less than " + fe.Param()
	case "gte":
		return "This field should be greater than " + fe.Param()
	case "min":
		return "A minimum of " + fe.Param() + " characters is required"
	case "email":
		return "Invalid email address"
	}

	return "Unkown Error"
}
