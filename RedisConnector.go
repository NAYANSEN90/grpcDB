package main

import (
	"fmt"
)

//RedisConfig defines the parameters required to make a redis server connection
type RedisConfig struct {
	host string
	port int
	user string
	pass string
}

//CreateRedisConnector function creates a connection to the redis server
func CreateRedisConnector() {

	fmt.Println("Redis connector created")
}
