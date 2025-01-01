package idgenerator

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateID generates a new MongoDB ObjectId.
func GenerateID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// IsValidID validates if a string is a valid MongoDB ObjectId.
func IsValidID(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}

func ToPersistenceID(id string) (primitive.ObjectID, error) {
	idResult, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return idResult, nil
}
