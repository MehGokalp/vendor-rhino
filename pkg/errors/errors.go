package errors

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// ValidationError will help to return customized Error info
// {"database": {"hello":"no such table", error: "not_exists"}}
type ValidationError struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewHttpValidationError handles the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewHttpValidationError(err error) ValidationError {
	res := ValidationError{}
	res.Errors = make(map[string]interface{})
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, v := range errs {
			// can translate each error one at a time.
			if v.Param() != "" {
				res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
			} else {
				res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
			}
		}

		return res
	}

	res.Errors["message"] = err.Error()

	return res
}
