package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureMongoIndexes(db *mongo.Database) {

	ctx := context.Background()

	// USERS COLLECTION
	_, err := db.Collection("users").Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"user_id": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("idx_user_user_id"),
		},
		{
			Keys: bson.M{"email": 1},
			Options: options.Index().
				SetUnique(true).
				SetName("idx_user_email"),
		},
		{
			Keys: bson.M{"role": 1},
			Options: options.Index().
				SetName("idx_user_role"),
		},
	})
	if err != nil {
		log.Fatal(" User indexes error:", err)
	}

	// PROJECTS COLLECTION
	_, err = db.Collection("projects").Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"ownerId": 1},
			Options: options.Index().
				SetName("idx_project_owner"),
		},
		{
			Keys: bson.M{"memberIds": 1},
			Options: options.Index().
				SetName("idx_project_members"),
		},
	})
	if err != nil {
		log.Fatal("Project indexes error:", err)
	}

	// TASKS COLLECTION
	_, err = db.Collection("tasks").Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"projectId": 1},
			Options: options.Index().
				SetName("idx_task_project"),
		},
		{
			Keys: bson.M{"assignedTo": 1},
			Options: options.Index().
				SetName("idx_task_assigned"),
		},
		{
			Keys: bson.M{"status": 1},
			Options: options.Index().
				SetName("idx_task_status"),
		},
	})
	if err != nil {
		log.Fatal("Task indexes error:", err)
	}

	log.Println(" MongoDB indexes ensured")
}
