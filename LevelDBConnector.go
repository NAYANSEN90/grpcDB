package main

import (
	"fmt"
)

//LevelDBConfig defines the parameteres required to make a level db connection
type LevelDBConfig struct {
	database string
}

//CreateConnector function creates a connection to the leveldb database
func CreateConnector() {
	fmt.Println("Created level db connector")
}
