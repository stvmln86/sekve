package test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestMockDB(t *testing.T) {
	// successs
	dbse := MockDB(t)
	assert.NotNil(t, dbse)

	// confirm - mock data
	dbse.View(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck := tx.Bucket([]byte(name))
			assert.NotNil(t, buck)

			for attr, want := range pairs {
				data := buck.Get([]byte(attr))
				assert.Equal(t, want, string(data))
			}
		}

		return nil
	})
}

func TestRequest(t *testing.T) {
	// success
	r := Request("GET", "/", "body", map[string]string{"attr": "data"})
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.String())

	// confirm - body
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)

	// confirm - path values
	data := r.PathValue("attr")
	assert.Equal(t, "data", data)
}
