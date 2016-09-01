// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package repos

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	d "org.desmax/gdgsardegna/domain"
)

const GdgBucket = "GDGBucket"

func ReadMsg(db bolt.DB, key string, logger log.Logger) []byte {
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(GdgBucket))
		variable := b.Get([]byte(key))
		v = variable
		fmt.Println(string(v))
		return nil
	})
	return v
}

func ReadAll(db bolt.DB, bucketName string, logger log.Logger) []d.Message {
	var msgs = make([]d.Message, 10)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var msg = d.Message{ID: int(binary.BigEndian.Uint64(k)), CONTENT: string(v)}
			msgs = append(msgs, msg)
		}
		return nil
	})
	return  msgs
}

func SaveOrUpdateMsg(db bolt.DB, msg d.Message, logger log.Logger) []byte {
	var v []byte
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(GdgBucket))
		if (msg.ID == 0) {
			id, errS := b.NextSequence()
			logger.Println(id)
			if errS != nil {
				logger.Println(errS)
			}
			msg.ID = int(id)
		}
		buf, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		var k = itob(msg.ID)
		err = b.Put(k, buf)
		if err != nil {
			logger.Printf("Error %s", err)
		}
		v = k
		return err
	})
	return v
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
