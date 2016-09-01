package repos

import (
	"testing"

	"github.com/boltdb/bolt"
	"time"
	"fmt"
	"org.desmax/gdgsardegna/domain"
	"log"
	"os"
	"encoding/json"
	"encoding/binary"
)

func TestUpdate(t *testing.T) {
	fmt.Println("TestUpdate")
	db, err := bolt.Open("./../../../../gdg-test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	msg := *new(domain.Message)
	msg.ID = 42
	msg.CONTENT = "Hello World !"

	logFile, _ := os.Create("./../../../../test.txt")
	logger := log.New(logFile, "gdg-test", log.Ldate|log.Lshortfile)
	var key []byte = SaveOrUpdateMsg(*db,msg, *logger)

	id := binary.BigEndian.Uint64(key)
	if(id == 0 ){
		t.Fail()
	}
	if(id != 42 ){
		t.Fail()
	}
}

func TestRead(t *testing.T) {
	fmt.Println("TestRead")
	db, err := bolt.Open("./../../../../gdg-test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	msg := *new(domain.Message)
	msg.ID = 13
	msg.CONTENT = "Hello World !"
	fmt.Println(msg)
	logFile, _ := os.Create("./../../../../test.txt")
	logger := log.New(logFile, "gdg-test", log.Ldate|log.Lshortfile)
	var result []byte = SaveOrUpdateMsg(*db,msg, *logger)
	fmt.Println(string(result))
	var msgByte []byte = ReadMsg(*db,string(result), *logger)
	var msgRead domain.Message
	err = json.Unmarshal(msgByte, &msgRead)
	if(err != nil || msgRead.ID == 0 || msgRead.CONTENT == ""){
		t.Fail()
	}
	logger.Println(msgRead)
}
