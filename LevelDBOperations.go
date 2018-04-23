package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"os"
)

var storage          *leveldb.DB    = nil

type LevelDBConfig struct {
	database string
}

func initLevelDBClient(config *Configuration) {
	options := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	database, err := leveldb.OpenFile(config.leveldb.database, options)
	if nil != err {
		database, err = leveldb.RecoverFile(config.leveldb.database, options)
	}
	if nil == err {
		storage = database
	} else {
		fmt.Println("Failed to open level db connection")
		os.Exit(-1)
	}
}


func closeLevelDBClient() {
	if storage != nil {
		storage.Close()
	}
}
