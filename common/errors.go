package common

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// CommonError will help to return customized Error info
// {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewValidatorError handles the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	if errs, ok := err.(validator.ValidationErrors); ok == true {
		for _, v := range errs {
			// can translate each error one at a time.
			//fmt.Println("gg",v.NameNamespace)
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

// NewError warps the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
