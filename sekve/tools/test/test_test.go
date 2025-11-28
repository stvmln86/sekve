package test

import (
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
