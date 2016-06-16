package db

import (
	"bytes"
	"fmt"
	"log"
	"path"
//	"time"

        "github.com/boltdb/bolt"
)

type DB struct {
	File	string
	Conn	*bolt.DB
}

type KeyValue struct {
	Key	string
	Value	string
}

type KV []KeyValue

func NewDB(appPath string) (*DB, error) {
	file := path.Join(appPath, "data.db")

	//db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	db, err := bolt.Open(file, 0600, nil)
	if err != nil {
		return &DB{}, err
	}

	return &DB{
		File: file,
		Conn: db,
	}, nil
}

func RecoverFunc() {
	if r := recover(); r != nil {
		log.Println("Recovered:", r)
	}
}

func (db *DB) List(b string) KV {
	elems := KV{}

	defer db.Conn.Close()

	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			elems = append(elems, KeyValue{Key: string(k), Value: string(v)})
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return elems
}

func (db *DB) Get(b, key string) []byte {
	value := []byte("")

	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))

		value = bucket.Get([]byte(key))

		return nil
	})
	if err != nil {
		panic(err)
	}

	return value
}

func (db *DB) Add(b, key string, value []byte) {
	defer db.Conn.Close()

	err := db.Conn.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (db *DB) Remove(b, key string) {
	defer db.Conn.Close()

	err := db.Conn.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b))
		if err != nil {
			return err
		}

		err = bucket.Delete([]byte(key))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (db *DB) KeyExists(b, key string) bool {
	ok := true

	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))

		v := bucket.Get([]byte(key))
		if v == nil {
			ok = false
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return ok
}

func (db *DB) KeyExistsRecursive(b, key string) bool {
	fmt.Println(key)

	ok := false

	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))

		v := bucket.Get([]byte(key))
		if v != nil {
			fmt.Println(v)
			if bytes.Equal([]byte("r"), v) {
				ok = true
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return ok
}

func (db *DB) Count(b string) int {
	count := 0

	err := db.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b))

		bs := bucket.Stats()
		if bs.KeyN > 0 {
			count = bs.KeyN
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return count
}
