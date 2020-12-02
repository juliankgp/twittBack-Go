package db

import (
	"context"
	"fmt"
	"time"

	"github.com/juliankgp/twittBack-Go/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckRelationship : Search another user with the same id
func CheckRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	col := MongoCN.Database("twittBack").Collection("relationship")

	condition := bson.M{
		"userId":          t.UserID,
		"userReferenceId": t.UserReferenceID,
	}

	var result models.Relationship
	err := col.FindOne(ctx, condition).Decode(&result)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
