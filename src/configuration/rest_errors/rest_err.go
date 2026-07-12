package resterrors

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int64   `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//necessario pra usar o tipo RestErr no tratamento nativo de erro em Go ex.: resp, err = request()
func (r *RestErr) Error() string{
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func newBadRequestErr(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func newBadRequestValidationErr(message string, causes []Cause) *RestErr{
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes: causes,
	}
}

func newInternalServerErr(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}


func newForbiddenErr(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err:     "forbidden_error",
		Code:    http.StatusForbidden,
	}
}


func newNotFoundErr(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err:     "not_found_error",
		Code:    http.StatusNotFound,
	}
}