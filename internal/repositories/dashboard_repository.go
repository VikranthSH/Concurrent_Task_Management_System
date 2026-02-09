package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DashboardRepository interface {
	GetAdminDashboard(ctx context.Context, adminID string) ([]bson.M, error)
}

type dashboardRepository struct {
	db *mongo.Database
}

func NewDashboardRepository(db *mongo.Database) DashboardRepository {
	return &dashboardRepository{db: db}
}

// ADMIN DASHBOARD (Aggregation)
func (r *dashboardRepository) GetAdminDashboard(
	ctx context.Context,
	adminID string,
) ([]bson.M, error) {

	adminObjID, err := primitive.ObjectIDFromHex(adminID)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{

		// Only employees
		{{Key: "$match", Value: bson.M{"role": "employee"}}},

		// Join projects OWNED by this admin
		{{Key: "$lookup", Value: bson.M{
			"from": "projects",
			"let": bson.M{"userId": "$_id"},
			"pipeline": bson.A{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{
							"$and": bson.A{
								bson.M{"$eq": bson.A{"$ownerId", adminObjID}},
								bson.M{"$in": bson.A{"$$userId", "$memberIds"}},
							},
						},
					},
				},
			},
			"as": "projects",
		}}},

		//  Remove users with NO projects under this admin
		{{Key: "$match", Value: bson.M{
			"projects.0": bson.M{"$exists": true},
		}}},

		// Join tasks
		{{Key: "$lookup", Value: bson.M{
			"from": "tasks",
			"localField": "_id",
			"foreignField": "assignedTo",
			"as": "tasks",
		}}},

		// Final shape
		{{Key: "$project", Value: bson.M{
			"_id": 1,
			"name": 1,
			"user_id": 1,
			"role": 1,
			"projects": 1,
			"tasks": 1,
		}}},
	}

	cursor, err := r.db.Collection("users").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
