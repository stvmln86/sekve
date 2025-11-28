// Package test implements unit testing data and functions.
package test

import (
	"io"
	"net/http/httptest"
)

// Response returns the status integer and body string from a ResponseRecorder.
func Response(w *httptest.ResponseRecorder) (int, string, error) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	return rslt.StatusCode, string(bytes), err
}
