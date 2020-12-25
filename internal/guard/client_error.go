package guard

import "errors"

type ClientError struct {
	Err error
}

func NewClientError(err error) error {
	return ClientError{Err: err}
}

func (e ClientError) Error() string {
	return "client error"
}

func (e ClientError) Unwrap() error {
	return e.Err
}

func IsClientError(err error) bool {
	return errors.As(err, new(ClientError))
}
