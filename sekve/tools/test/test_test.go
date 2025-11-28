package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, `"body"`)

	// success
	AssertJSON(t, w, http.StatusOK, "body")
}

func TestRequest(t *testing.T) {
	// success
	r := Request("GET", "/", "body")
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.Path)

	// confirm - body
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}
