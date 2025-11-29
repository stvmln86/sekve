// Package bolt implements Bolt database handling functions.
package bolt

import (
	"strings"

	"go.etcd.io/bbolt"
)

// Delete deletes an existing database bucket.
func Delete(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// Exists returns true if a database bucket exists.
func Exists(db *bbolt.DB, name string) (bool, error) {
	var okay bool
	return okay, db.View(func(tx *bbolt.Tx) error {
		okay = tx.Bucket([]byte(name)) != nil
		return nil
	})
}

// Read returns an existing database bucket as a string map.
func Read(db *bbolt.DB, name string) (map[string]string, error) {
	var pairs map[string]string
	return pairs, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			pairs = make(map[string]string)
			return buck.ForEach(func(attr, data []byte) error {
				pairs[string(attr)] = string(data)
				return nil
			})
		}

		return nil
	})
}

// Search returns all existing database bucket names matching a prefix string.
func Search(db *bbolt.DB, pref string) ([]string, error) {
	var names []string
	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			if strings.HasPrefix(string(name), pref) {
				names = append(names, string(name))
			}

			return nil
		})
	})
}

// Write writes a new or existing database bucket with a string map.
func Write(db *bbolt.DB, name string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, data := range pairs {
			if err := buck.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
