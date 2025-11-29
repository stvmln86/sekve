package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sekve/sekve/tools/test"
	"go.etcd.io/bbolt"
)

func TestDelete(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Delete(db, "user.0000000000000000000000")
	assert.NoError(t, err)

	// confirm - bucket deleted
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("user.0000000000000000000000"))
		assert.Nil(t, buck)
		return nil
	})
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - true
	okay, err := Exists(db, "user.0000000000000000000000")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Exists(db, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	pairs, err := Read(db, "user.0000000000000000000000")
	assert.Equal(t, map[string]string{
		"addr": "1.2.3.4",
		"init": "1000",
	}, pairs)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	names, err := Search(db, "pair.0000000000000000000000")
	assert.Equal(t, []string{
		"pair.0000000000000000000000.alpha",
		"pair.0000000000000000000000.bravo",
	}, names)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Write(db, "name", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// confirm - bucket created
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("name"))
		assert.NotNil(t, buck)

		data := buck.Get([]byte("attr"))
		assert.Equal(t, []byte("data"), data)
		return nil
	})
}
