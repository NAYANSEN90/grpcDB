package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var redisConnection  *redis.Client  = nil

//RedisConfig defines the parameters required to make a redis server connection
type RedisConfig struct {
	host string
	port int
	user string
	pass string
}


func callGet(key string) (string, error) {

	if nil != redisConnection {
		resp,err := redisConnection.Get(key).Result()
		if nil != err {
			return resp, nil
		}
	}
	return  "", fmt.Errorf("not found")
}


func initRedisClient(config * Configuration){
	var address string
	address  = fmt.Sprintf("%s:%d" , config.redis.host , config.redis.port)
	fmt.Printf("Connecting to redis server %s\n" , address)

	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Failed to establish connection with redis server")
		os.Exit(-1)
	} else {
		fmt.Printf("Received pong from server : %s \n", pong)
	}

	redisConnection = client
}

func closeRedisDBClient() {
	if redisConnection != nil {
		redisConnection.Close()
	}
}