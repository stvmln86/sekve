package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprint(w, "body")

	// success
	code, body, err := Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "body", body)
	assert.NoError(t, err)
}
