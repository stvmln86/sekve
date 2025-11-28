// Package http implements HTTP protocol functions.
package http

import (
	"encoding/json"
	"net/http"
	"strings"
)

// writeJSON writes a marshalled JSON object to a ResponseWriter.
func writeJSON(w http.ResponseWriter, code int, data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		code = http.StatusInternalServerError
		bytes = []byte(`{"status": "error", "message": "internal server error"}`)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(bytes)
}

// Error writes a JSend error message to a ResponseWriter.
func Error(w http.ResponseWriter, code int, text string) {
	writeJSON(w, code, map[string]any{
		"status":  "error",
		"message": text,
	})
}

// ErrorCode writes a JSend error with a default message to a ResponseWriter.
func ErrorCode(w http.ResponseWriter, code int) {
	text := http.StatusText(code)
	Error(w, code, strings.ToLower(text))
}

// Fail writes a JSend failure payload to a ResponseWriter.
func Fail(w http.ResponseWriter, code int, pairs map[string]string) {
	writeJSON(w, code, map[string]any{
		"status": "fail",
		"data":   pairs,
	})
}

// Success writes a JSend success payload to a ResponseWriter.
func Success(w http.ResponseWriter, code int, data any) {
	writeJSON(w, code, map[string]any{
		"status": "success",
		"data":   data,
	})
}

// Unmarshal returns an unmarshalled JSON object from a Request body.
func Unmarshal(r *http.Request) (map[string]any, error) {
	var pairs map[string]any
	return pairs, json.NewDecoder(r.Body).Decode(&pairs)
}
