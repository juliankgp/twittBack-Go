package models

// Relationship model to save users relationships
type Relationship struct {
	UserID          string `bson:"userId" json:"userId"`
	UserReferenceID string `bson:"userReferenceId" json:"userReferenceId"`
}
