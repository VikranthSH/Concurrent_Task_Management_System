package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	Priority    string             `bson:"priority" json:"priority"`
	DueDate     time.Time          `bson:"dueDate" json:"dueDate"`
	ProjectID   primitive.ObjectID `bson:"projectId" json:"projectId"`
	AssignedTo  primitive.ObjectID `bson:"assignedTo" json:"assignedTo"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
