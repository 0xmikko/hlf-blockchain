/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

func HttpBadRequestError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Wrong request",
		Code:    http.StatusBadRequest,
		Err:     e,
	}
}

func HttpForbiddenRequestError() ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Forbidden",
		Code:    http.StatusBadRequest,
		Err:     errors.New("Forbidden"),
	}
}

func UnknownError(e error) ApiError {
	_, file, line, _ := runtime.Caller(2)
	return ApiError{
		Module:  fmt.Sprintf("%s at %d", file, line),
		Message: "Unknown error",
		Code:    http.StatusInternalServerError,
		Err:     e,
	}
}
