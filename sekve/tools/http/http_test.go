package http

import (
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sekve/sekve/tools/test"
)

func TestWriteJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success - marshal succeeded
	writeJSON(w, http.StatusOK, "data")
	test.AssertJSON(t, w, http.StatusOK, "data")

	// setup
	w = httptest.NewRecorder()

	// success - marshal failed
	writeJSON(w, http.StatusOK, math.NaN)
	test.AssertJSON(t, w, http.StatusInternalServerError, map[string]any{
		"status":  "error",
		"message": "internal server error",
	})
}

func TestError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Error(w, http.StatusInternalServerError, "text")
	test.AssertJSON(t, w, http.StatusInternalServerError, map[string]any{
		"status":  "error",
		"message": "text",
	})
}

func TestErrorCode(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	ErrorCode(w, http.StatusInternalServerError)
	test.AssertJSON(t, w, http.StatusInternalServerError, map[string]any{
		"status":  "error",
		"message": "internal server error",
	})
}

func TestFail(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Fail(w, http.StatusNotFound, map[string]string{"key": "value"})
	test.AssertJSON(t, w, http.StatusNotFound, map[string]any{
		"status": "fail",
		"data":   map[string]any{"key": "value"},
	})
}

func TestSuccess(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Success(w, http.StatusOK, "data")
	test.AssertJSON(t, w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   "data",
	})
}

func TestUnmarshal(t *testing.T) {
	// setup
	r := test.Request("GET", "/", `{"key": "value"}`)

	// success
	pairs, err := Unmarshal(r)
	assert.Equal(t, map[string]any{"key": "value"}, pairs)
	assert.NoError(t, err)
}
