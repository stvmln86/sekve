// Package test implements unit testing data and functions.
package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]string{
	"user.aaaa": {
		"addr": "1.2.3.4",
		"init": "1000",
	},

	"pair.aaaa.alpha": {
		"body": "Alpha value.",
		"hash": "49e8c3bb0a4c0773b54af4aee638ef128c5dceae19b2e5adba57f0bdc33d4840",
		"init": "2000",
	},

	"pair.aaaa.bravo": {
		"body": "Bravo value.",
		"hash": "e628d55d2c5c5e47bda1fbb4fe8c8a365eb12c89d2745346216e20cad0b4a0c3",
		"init": "3000",
	},
}

// MockDB returns a temporary empty Bolt database.
func MockDB(t *testing.T) *bbolt.DB {
	dest := filepath.Join(t.TempDir(), "bolt.dbse")
	dbse, _ := bbolt.Open(dest, 0600, nil)
	dbse.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck, _ := tx.CreateBucketIfNotExists([]byte(name))
			for attr, data := range pairs {
				buck.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return dbse
}

// Request returns a new mock Request with a body string and value pairs.
func Request(meth, path, body string, pairs map[string]string) *http.Request {
	rqst := httptest.NewRequest(meth, path, strings.NewReader(body))
	for attr, data := range pairs {
		rqst.SetPathValue(attr, data)
	}

	return rqst
}
