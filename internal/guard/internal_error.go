package guard

import "errors"

type InternalError struct {
	Err error
}

func NewInternalError(err error) error {
	return InternalError{Err: err}
}

func (e InternalError) Error() string {
	return "internal error"
}

func (e InternalError) Unwrap() error {
	return e.Err
}

func IsInternalError(err error) bool {
	return errors.As(err, new(InternalError))
}
