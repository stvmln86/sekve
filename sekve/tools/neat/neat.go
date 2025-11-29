// Package neat implements value sanitisation functions.
package neat

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// trim returns a string trimmed to a maximum length.
func trim(text string, size int) string {
	if len(text) > size {
		return text[:size]
	}

	return text
}

// Body returns a whitespace-trimmed pair body string.
func Body(body string, size int) string {
	body = strings.TrimSpace(body)
	return trim(body, size)
}

// Hash returns a SHA256 hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}

// Join returns a dot-joined database string from one or more strings.
func Join(form string, elems ...string) string {
	join := strings.Join(elems, ".")
	return fmt.Sprintf("%s.%s", form, join)
}

// Name returns a lowercase pair name string.
func Name(name string, size int) string {
	name = strings.TrimSpace(name)
	name = trim(name, size)
	return strings.ToLower(name)
}

// Time returns a Time object from a Unix UTC string.
func Time(unix string) time.Time {
	unix = strings.TrimSpace(unix)
	uint, _ := strconv.ParseInt(unix, 10, 64)
	return time.Unix(uint, 0).Local()
}

// UUID returns a new or existing lowercase UUID string.
func UUID(uuids ...string) string {
	if len(uuids) == 0 {
		uuid := uuid.New().String()
		return strings.ReplaceAll(uuid, "-", "")
	}

	uuid := strings.TrimSpace(uuids[0])
	uuid = trim(uuid, 32)
	return strings.ToLower(uuid)
}

// Unix returns a Unix UTC string from a Time object.
func Unix(tobj time.Time) string {
	uint := tobj.Unix()
	return strconv.FormatInt(uint, 10)
}
