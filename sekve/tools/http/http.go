// Package http implements HTTP request and response functions.
package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Read returns a Request's raw body as a string.
func Read(r *http.Request) string {
	bytes, _ := io.ReadAll(r.Body)
	return string(bytes)
}

// Write writes a formatted plaintext string to a ResponseWriter.
func Write(w http.ResponseWriter, code int, body string, elems ...any) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, body, elems...)
}

// WriteError writes a formatted plaintext error message to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, body string, elems ...any) {
	body = fmt.Sprintf(body, elems...)
	Write(w, code, "error %d: %s", code, body)
}

// WriteErrorCode writes a static plaintext error code to a ResponseWriter.
func WriteErrorCode(w http.ResponseWriter, code int) {
	body := http.StatusText(code)
	body = strings.ToLower(body)
	Write(w, code, "error %d: %s", code, body)
}
