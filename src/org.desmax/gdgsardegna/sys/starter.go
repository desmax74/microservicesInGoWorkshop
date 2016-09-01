// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package sys

import (
	"github.com/boltdb/bolt"
	"log"
	"net"
	"org.desmax/gdgsardegna/config"
	"time"
	"os"
	"fmt"
)

func StartLogServer(conf config.Configuration) net.Conn {
	logServer, err := net.Dial("tcp", conf.LogServer)
	if err != nil {
		panic("Failed to connect to " + conf.LogServer)
	}
	return logServer
}

func StartLogger(conf config.Configuration, logServer net.Conn) log.Logger {
	logger := log.New(logServer, "gdg ", log.Ldate|log.Lshortfile)
	return *logger
}

func StartLoggerFS(conf config.Configuration) (log.Logger, *os.File)  {
	logFile, errLog := os.Create(conf.LogPath)
	if(errLog != nil){
		fmt.Println(errLog)
	}
	logger := log.New(logFile, "gdg ", log.Ldate|log.Lshortfile)
	return  *logger, logFile
}

func StartLoggerStdOut(conf config.Configuration, level string) (log.Logger)  {
	logger := log.New(os.Stdout, level, log.Ldate|log.Lshortfile)
	return  *logger
}

func StartDB(conf config.Configuration, logger log.Logger) bolt.DB {
	db, err := bolt.Open("gdg.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil{
		logger.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("GDGBucket"))
		if err != nil {
			logger.Printf("create bucket:", err)
		} else {
			logger.Println("created bucket: GDGBucket")
		}
		return nil
	})
	return *db
}
