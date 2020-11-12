package db

import (
	"context"
	"time"

	"github.com/juliankgp/twittBack-Go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRecord : Func to create a user
func InsertRecord(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittBack")
	col := db.Collection("users")

	u.Password, _ = EncryptPass(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
