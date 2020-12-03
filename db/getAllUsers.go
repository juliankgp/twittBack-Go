package db

import (
	"context"
	"fmt"
	"time"

	"github.com/juliankgp/twittBack-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUsers : get all user with conditions
func GetAllUsers(ID string, page int64, search string, kind string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := MongoCN.Database("twittBack").Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relationship
		r.UserID = ID
		r.UserReferenceID = s.ID.Hex()

		include = false

		found, err = CheckRelationship(r)
		if kind == "new" && found == false {
			include = true
		}
		if kind == "follow" && found == true {
			include = true
		}
		if r.UserReferenceID == ID {
			include = false
		}
		if include {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true

}
