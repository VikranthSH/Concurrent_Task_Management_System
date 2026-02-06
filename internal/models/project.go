package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	OwnerID     primitive.ObjectID   `bson:"ownerId" json:"ownerId"`
	MemberIDs   []primitive.ObjectID `bson:"memberIds" json:"memberIds"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
}
