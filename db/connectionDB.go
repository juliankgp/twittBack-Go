package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN parameter to connection to bd
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://juliankgp:ivqymgvmc@twittbackgo.3wlwe.mongodb.net/test")

// ConnectDB Function to connect to the BD
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conecction Successfully to DB")
	return client

}

// CheckConnection Ping to BD
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
