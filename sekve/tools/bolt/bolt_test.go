package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sekve/sekve/tools/test"
	"go.etcd.io/bbolt"
)

func TestDelete(t *testing.T) {
	// setup
	dbse := test.MockDB(t)

	// success
	err := Delete(dbse, "user.aaaa")
	assert.NoError(t, err)

	// confirm - bucket deleted
	dbse.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("user.aaaa"))
		assert.Nil(t, buck)
		return nil
	})
}

func TestExists(t *testing.T) {
	// setup
	dbse := test.MockDB(t)

	// success - true
	okay, err := Exists(dbse, "user.aaaa")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Exists(dbse, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	// setup
	dbse := test.MockDB(t)

	// success
	pairs, err := Read(dbse, "user.aaaa")
	assert.Equal(t, map[string]string{
		"addr": "1.2.3.4",
		"init": "1000",
	}, pairs)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	dbse := test.MockDB(t)

	// success
	names, err := Search(dbse, "pair.aaaa")
	assert.Equal(t, []string{
		"pair.aaaa.alpha",
		"pair.aaaa.bravo",
	}, names)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	dbse := test.MockDB(t)

	// success
	err := Write(dbse, "name", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// confirm - bucket created
	dbse.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("name"))
		assert.NotNil(t, buck)

		data := buck.Get([]byte("attr"))
		assert.Equal(t, []byte("data"), data)
		return nil
	})
}
