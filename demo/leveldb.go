package main

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	db, err := leveldb.OpenFile("./local.db", nil)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer db.Close()

	for i := 0; i < 100000; i++ {
		err = db.Put([]byte("hello"+gconv.String(i)), []byte("world"+gconv.String(i)), nil)
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		log.Println(gconv.String(iter.Key()) + gconv.String(iter.Value()))
	}

}
