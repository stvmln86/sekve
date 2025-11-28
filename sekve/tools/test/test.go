// Package test implements unit testing data and functions.
package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertJSON asserts a ResponseRecorder's JSON body is equal to a value.
func AssertJSON(t *testing.T, w *httptest.ResponseRecorder, code int, want any) {
	var data any
	rslt := w.Result()
	err := json.NewDecoder(rslt.Body).Decode(&data)
	assert.Equal(t, code, rslt.StatusCode)
	assert.Equal(t, want, data)
	assert.NoError(t, err)
}

// Request returns a new mock Request with a body string and path values.
func Request(meth, path, body string) *http.Request {
	buff := strings.NewReader(body)
	return httptest.NewRequest(meth, path, buff)
}
