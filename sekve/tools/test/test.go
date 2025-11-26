// Package test implements unit testing data and functions.
package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

// Request returns a new Request with a body string and value pairs.
func Request(mthd, path, body string, pairs map[string]string) *http.Request {
	r := httptest.NewRequest(mthd, path, strings.NewReader(body))

	for name, data := range pairs {
		r.SetPathValue(name, data)
	}

	return r
}

// Response returns the status integer and body string from a ResponseRecorder.
func Response(w *httptest.ResponseRecorder) (int, string, error) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	return rslt.StatusCode, string(bytes), err
}
