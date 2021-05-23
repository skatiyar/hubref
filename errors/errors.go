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
)

type Error struct {
	Code     int
	Message  string
	RawError error
}

func (e Error) Error() string {
	return e.RawError.Error()
}

func InvalidDataError(text string) error {
	return Error{
		Code:     400,
		Message:  "Invalid Data",
		RawError: errors.New(text),
	}
}

func DBError(text string) error {
	return Error{
		Code:     500,
		Message:  "DB Error",
		RawError: errors.New(text),
	}
}

func InternalServerError(text string) error {
	return Error{
		Code:     500,
		Message:  "Internal Server Error",
		RawError: errors.New(text),
	}
}

func NotFoundError(text string) error {
	return Error{
		Code:     404,
		Message:  "Record Not Found",
		RawError: errors.New(text),
	}
}
