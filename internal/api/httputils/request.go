package httputils

import (
	"net/http"
	"strconv"
)

func GetUserName(r *http.Request) string {
	return r.Header.Get("x-user-name")
}

func GetUint16FromQuery(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint16(v), err
}
