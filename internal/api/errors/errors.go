package errors

import "fmt"

type WrongQueryValueError struct {
	t, name, value string
}

func (w WrongQueryValueError) Error() string {
	return fmt.Sprintf("query %s should be valid %s, but provided '%s'",
		w.name, w.t, w.value)
}

func NewWrongQueryValueError(queryType, queryName, queryValue string) error {
	return WrongQueryValueError{
		t:     queryType,
		name:  queryName,
		value: queryValue,
	}
}
