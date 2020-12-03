package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FollowerTweets struc to return Follower Tweets
type FollowerTweets struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID          string             `bson:"userId" json:"userId,omitempty"`
	UserReferenceID string             `bson:"userReferenceId" json:"userReferenceId,omitempty"`
	Tweet           struct {
		Mensaje string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
