package classes

import (
	"github.com/boltdb/bolt"
	"github.com/dkadio/dysr/util"
	"log"
)

type boltAdapter struct {
	storeName  string
	bucketName string
}

func NewBoltAdapter(c util.Config) boltAdapter {
	db := openDataBase(c.StoreName)
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(c.BucketName))
		if err != nil {
			panic(err)
		}
		return nil
	})
	defer db.Close()

	return boltAdapter{c.StoreName, c.BucketName}
}

//abstract those functions threw interface
func (b boltAdapter) putValueFor(key, value string) error {
	db := openDataBase(b.storeName)

	defer db.Close()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

//abstract those functions threw interface
func openDataBase(storeName string) *bolt.DB {
	db, err := bolt.Open(storeName, 0600, nil)
	if err != nil {
		panic(err)
	}
	return db
}

//abstract those functions threw interface
func (b boltAdapter) getValueFor(key string) string {
	db := openDataBase(b.storeName)

	defer db.Close()
	var v string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		v = string(b.Get([]byte(key)))
		return nil
	})
	if v == "" {
		log.Println("No Value for key", key)
	} else {
		log.Println("found key/value", key, "=", v)
	}
	return v
}
