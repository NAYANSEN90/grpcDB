package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var storage *leveldb.DB = nil
var defaultConfig *Configuration = nil

//go:generate protoc -I proto/ proto/server.proto --go_out=plugins=grpc:protogo

func init() {
	options := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	database, err := leveldb.OpenFile("levelR.db", options)
	if nil != err {
		database, err = leveldb.RecoverFile("levelR.db", options)
	}
	if nil == err {
		storage = database
	} else {
		os.Exit(-1)
	}
}

func closeDB() {
	if storage != nil {
		storage.Close()
	}
}

//Configuration holds the configuration for redis, leveldb and grpc server
type Configuration struct {
	redis   RedisConfig
	leveldb LevelDBConfig
	grpc    GrpcConfig
}

//TODO: Fix json marshalling
func parseConfig() Configuration {
	file := "./configuration.json"
	raw, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config Configuration
	json.Unmarshal(raw, &config)
	return config
}

func main() {

	defaultConfig = &(Configuration{
		RedisConfig{host: "127.0.0.1", port: 6379},
		LevelDBConfig{database: "levelR.db"},
		GrpcConfig{host: "127.0.0.1", port: 9999}})

	fmt.Println(defaultConfig.redis.host)
}
