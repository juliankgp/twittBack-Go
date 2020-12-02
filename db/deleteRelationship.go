package db

import (
	"context"
	"time"

	"github.com/juliankgp/twittBack-Go/models"
)

// DeleteRelationship Break relationship
func DeleteRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twittBack").Collection("relationship")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
