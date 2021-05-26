/*
 * NOTE
 * This package exposes an error struct
 * which satifies error interface. Any error
 * in the code should be wrapped using one of the
 * following functions before returning to caller
 * function. This simplifies error code handling
 * on API level.
 * TODO
 * Is adding file and line number,
 * for stack trace too much ?
 */
package errors

import (
	"errors"
	"net/http"
)

type Error struct {
	Code     int
	Class    string
	RawError error
}

func (e Error) Error() string {
	return e.RawError.Error()
}

func InvalidDataError(text string) error {
	return Error{
		Code:     http.StatusBadRequest,
		Class:    "Invalid Data",
		RawError: errors.New(text),
	}
}

func DBError(text string) error {
	return Error{
		Code:     http.StatusInternalServerError,
		Class:    "DB Error",
		RawError: errors.New(text),
	}
}

func InternalServerError(text string) error {
	return Error{
		Code:     http.StatusInternalServerError,
		Class:    "Internal Server Error",
		RawError: errors.New(text),
	}
}

func NotFoundError(text string) error {
	return Error{
		Code:     http.StatusNotFound,
		Class:    "Record Not Found",
		RawError: errors.New(text),
	}
}
