package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"os"
)

var storage *leveldb.DB = nil

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

func callGet(key string) (string, error) {

	if nil != storage {
		val, err := storage.Get([]byte(key), nil)
		if nil != err {
			return "", err
		}
		return string(val), nil
	}
	return "", fmt.Errorf("db not found")
}

func callSet(key string, value string) error {
	if nil != storage {
		err := storage.Put([]byte(key), []byte(value), &opt.WriteOptions{Sync: true})
		if nil != err {
			return err
		}
		return nil
	}
	return fmt.Errorf("db not found")
}

func closeLevelDBClient() {
	if storage != nil {
		storage.Close()
	}
}
