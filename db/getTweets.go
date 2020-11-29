package db

import (
	"context"
	"log"
	"time"

	"github.com/juliankgp/twittBack-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTweets return all tweets with pagination
func GetTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twittBack").Collection("tweet")

	var results []*models.ReturnTweets

	condition := bson.M{
		"userId": ID,
	}

	params := options.Find()

	params.SetLimit(20)
	params.SetSort(bson.D{{Key: "date", Value: -1}})
	params.SetSkip((page - 1) * 20)

	dataColl, err := col.Find(ctx, condition, params)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for dataColl.Next(context.TODO()) {
		var registry models.ReturnTweets
		err := dataColl.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}
	return results, true

}
