package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var defaultConfig    *Configuration = nil

//go:generate protoc -I proto/ proto/server.proto --go_out=plugins=grpc:protogo


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
	fmt.Println(config)

	return config
}

func main() {

	defaultConfig = &(Configuration{
		RedisConfig{host: "127.0.0.1", port: 6379},
		LevelDBConfig{database: "levelR.db"},
		GrpcConfig{host: "127.0.0.1", port: 9999}})

	initRedisClient(defaultConfig)
	initLevelDBClient(defaultConfig)
	initGrpcServer(defaultConfig)

	defer closeRedisDBClient()
	defer closeLevelDBClient()
	defer shutdownGrpcServer()
}
