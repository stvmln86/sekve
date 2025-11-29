package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sekve/sekve/tools/test"
)

func TestRead(t *testing.T) {
	// setup
	rqst := test.Request("GET", "/", "body", nil)

	// success
	body := Read(rqst)
	assert.Equal(t, "body", body)
}

func TestWrite(t *testing.T) {
	// setup
	wrec := httptest.NewRecorder()

	// success
	Write(wrec, http.StatusOK, "%s", "body")

	// confirm - response
	assert.Equal(t, http.StatusOK, wrec.Code)
	assert.Equal(t, "text/plain; charset=utf-8", wrec.Header().Get("Content-Type"))
	assert.Equal(t, "body", wrec.Body.String())
}

func TestWriteError(t *testing.T) {
	// setup
	wrec := httptest.NewRecorder()

	// success
	WriteError(wrec, http.StatusNotFound, "%s", "body")

	// confirm - response
	assert.Equal(t, http.StatusNotFound, wrec.Code)
	assert.Equal(t, "error 404: body", wrec.Body.String())
}

func TestWriteErrorCode(t *testing.T) {
	// setup
	wrec := httptest.NewRecorder()

	// success
	WriteErrorCode(wrec, http.StatusNotFound)

	// confirm - response
	assert.Equal(t, http.StatusNotFound, wrec.Code)
	assert.Equal(t, "error 404: not found", wrec.Body.String())
}
