package repositories

import (
	"context"

	"Concurrent_Task_Management_System/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	FindAll(ctx context.Context) ([]models.Project, error)
	FindByOwnerID(ctx context.Context, ownerID primitive.ObjectID) ([]models.Project, error)
	FindByMemberID(ctx context.Context, userID primitive.ObjectID) ([]models.Project, error)

	UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type projectRepository struct {
	collection *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) ProjectRepository {
	return &projectRepository{
		collection: db.Collection("projects"),
	}
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) error {
	result, err := r.collection.InsertOne(ctx, project)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		project.ID = oid
	}

	return nil
}


func (r *projectRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	var project models.Project
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) FindAll(ctx context.Context) ([]models.Project, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Project, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"ownerId": userID},
			{"memberIds": userID},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
	)
	return err
}

func (r *projectRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
// OWNER ONLY
func (r *projectRepository) FindByOwnerID(
	ctx context.Context,
	ownerID primitive.ObjectID,
) ([]models.Project, error) {

	filter := bson.M{"ownerId": ownerID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// MEMBER ONLY
func (r *projectRepository) FindByMemberID(
	ctx context.Context,
	userID primitive.ObjectID,
) ([]models.Project, error) {

	filter := bson.M{"memberIds": userID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
