//Cannot encoding the bigcache object
package bolt

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
	"github.com/boltdb/bolt"
	"utils/file"
	"github.com/allegro/bigcache"
)

const (
	//---------JOBS PERSISTENCE CONSTANTS -----------//
	DB_DIR     string = "storage/"
	BUCKET_ID 	  = "mybucket"
	DB_NAME           = "data.db"
	READWRITE         = 0600
	READ              = 0666
	TIMEOUT           = 1
)

var (
	db *bolt.DB = new(bolt.DB)
)

func init() {
	tools.CreateDir(DB_DIR)
	var err error
	db, err = bolt.Open(DB_DIR + DB_NAME, READWRITE, &bolt.Options{Timeout: TIMEOUT * time.Second})
	if err != nil {
		log.Printf("-- cannot create %s \n", DB_NAME)
	}
}


//Delete the data using key
func DeleteJob(key string) error {
	// Delete the element using it key
	err := db.Update(
		func(tx *bolt.Tx) error {
			return tx.Bucket([]byte(BUCKET_ID)).Delete([]byte(key))
		})
	if err != nil {
		log.Printf("-- remove job %s failed\n", key)
		return err
	}
	return nil
}


// Save to Bolt
func Put(key string, data *bigcache.BigCache) error {
	//func Put(key string, data *interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BUCKET_ID))
		if err != nil {
			log.Println(err.Error())
			return err
		}

		//encode data before put in DB
		buffer := new(bytes.Buffer)
		enc := gob.NewEncoder(buffer)
		err = enc.Encode(data)
		if err != nil {
			log.Println("-- encoding data error", err)
			return err
		}

		return b.Put([]byte(key), buffer.Bytes())
	})
}

//Get Job from Bolt
func Get(key string) (data *bigcache.BigCache, err error) {
	//func Get(key string) (data *interface{}, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_ID))

		//Bucket exist??
		if b == nil {
			return nil
		}

		val := b.Get([]byte(key))
		dec := gob.NewDecoder(bytes.NewReader(val))
		err = dec.Decode(&data)
		return err
	})

	if err != nil {
		return nil, err
	}
	return data, err
}