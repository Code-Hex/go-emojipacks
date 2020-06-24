package slack

import "errors"

var (
	ErrReadLimitExceeded = errors.New("read limit exceeded")
	ErrTokenNotFound     = errors.New("token not found")
	ErrSomeThingWrong    = errors.New("something wrong (It might be page design has been changed)")
)

// OpError is the error type usually returned by functions in the loginslack
// package.
type OpError struct {
	// Op is the operation which caused the error, such as
	// "basic authentication" or "two factor authentication".
	Op string

	// Err is the error that occurred during the operation.
	Err error
}

func (e *OpError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Op + ": " + e.Err.Error()
}

// Unwrap unpacks wrapped error that occurred during the operation.
func (e *OpError) Unwrap() error { return e.Err }
