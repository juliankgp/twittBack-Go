package models

// Tweet Model to create tweets
type Tweet struct {
	Message string `bson:"message" json:"message,omitempty"`
}
