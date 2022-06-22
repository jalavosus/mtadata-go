package models

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/jalavosus/mtadata/models/apimethods"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

const (
	internalErrorMsg      string = "internal error"
	missingParamsErrorMsg string = "missing request parameters"
)

type ApiError interface {
	Proto() *protosv1.Error
	Error() string
	Cause() error
	Method() string
	Code() protosv1.ErrorCode

	baseMessage() string
}

type apiError struct {
	cause   error
	code    protosv1.ErrorCode
	msg     string
	baseMsg string
	method  apimethods.ApiMethod
}

func (e *apiError) Proto() *protosv1.Error {
	return &protosv1.Error{
		Code:    e.code,
		Message: apiErrorMsg(e),
		Method:  e.method.String(),
	}
}

func (e *apiError) Error() string {
	return e.msg
}

func (e *apiError) Cause() error {
	return e.cause
}

func (e *apiError) Method() string {
	return e.method.String()
}

func (e *apiError) Code() protosv1.ErrorCode {
	return e.code
}

func (e *apiError) baseMessage() string {
	return e.baseMsg
}

func InternalError(cause error, method apimethods.ApiMethod) ApiError {
	err := &apiError{
		cause:   cause,
		code:    protosv1.ErrorCode_InternalServerError,
		baseMsg: internalErrorMsg,
		method:  method,
	}

	err.msg = codedApiErrorMsg(err)

	return err
}

func EntityNotFoundError(entityName, entityId string, method apimethods.ApiMethod) ApiError {
	err := &apiError{
		cause:   errors.Errorf("%[1]s with id %[2]s not found", entityName, entityId),
		code:    protosv1.ErrorCode_EntityNotFound,
		baseMsg: "",
		method:  method,
	}

	err.msg = codedApiErrorMsg(err)

	return err
}

func MissingParametersError(method apimethods.ApiMethod, params ...string) ApiError {
	err := &apiError{
		cause:   errors.Errorf("%[1]s not provided", strings.Join(params, ", ")),
		code:    protosv1.ErrorCode_MissingParameters,
		baseMsg: missingParamsErrorMsg,
		method:  method,
	}

	err.msg = codedApiErrorMsg(err)

	return err
}

func apiErrorMsg(err ApiError) string {
	if err.baseMessage() == "" {
		return err.Cause().Error()
	}

	return fmt.Sprintf(
		"%[1]s: %[2]s",
		err.baseMessage(),
		err.Cause().Error(),
	)
}

func codedApiErrorMsg(err ApiError) string {
	if err.baseMessage() == "" {
		return fmt.Sprintf(
			"[%[1]d] %[2]s",
			err.Code(),
			err.Cause().Error(),
		)
	}

	return fmt.Sprintf(
		"[%[1]d] %[2]s: %[3]s",
		err.Code(),
		err.baseMessage(),
		err.Cause().Error(),
	)
}
