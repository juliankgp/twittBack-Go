package main

import (
	"log"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("Error connecting to the DB")
		return
	}
	handlers.Controller()
}
