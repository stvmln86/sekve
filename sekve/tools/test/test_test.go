package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssertJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, `"body"`)

	// success
	AssertJSON(t, w, http.StatusOK, "body")
}
