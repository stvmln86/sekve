package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// setup
	buff := strings.NewReader("body")
	rqst := httptest.NewRequest("GET", "/", buff)

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
	rslt := wrec.Result()
	bytes, err := io.ReadAll(rslt.Body)
	assert.Equal(t, http.StatusOK, rslt.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", rslt.Header.Get("Content-Type"))
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}

func TestWriteError(t *testing.T) {
	// setup
	wrec := httptest.NewRecorder()

	// success
	WriteError(wrec, http.StatusNotFound, "%s", "body")

	// confirm - response
	rslt := wrec.Result()
	bytes, err := io.ReadAll(rslt.Body)
	assert.Equal(t, http.StatusNotFound, rslt.StatusCode)
	assert.Equal(t, "error 404: body", string(bytes))
	assert.NoError(t, err)
}

func TestWriteErrorCode(t *testing.T) {
	// setup
	wrec := httptest.NewRecorder()

	// success
	WriteErrorCode(wrec, http.StatusNotFound)

	// confirm - response
	rslt := wrec.Result()
	bytes, err := io.ReadAll(rslt.Body)
	assert.Equal(t, http.StatusNotFound, rslt.StatusCode)
	assert.Equal(t, "error 404: not found", string(bytes))
	assert.NoError(t, err)
}
