package repositories

import (
	"context"

	"Concurrent_Task_Management_System/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error)
	FindAll(ctx context.Context) ([]models.Task, error)
	FindByProjectID(ctx context.Context, projectID primitive.ObjectID) ([]models.Task, error)
	FindByAssignedUser(ctx context.Context, userID primitive.ObjectID) ([]models.Task, error)
	FindByStatus(ctx context.Context, status string) ([]models.Task, error)
	UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
	}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid
	}

	return nil
}


func (r *taskRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
	var task models.Task
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) FindAll(ctx context.Context) ([]models.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindByProjectID(ctx context.Context, projectID primitive.ObjectID) ([]models.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"projectId": projectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindByAssignedUser(ctx context.Context, userID primitive.ObjectID) ([]models.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"assignedTo": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindByStatus(ctx context.Context, status string) ([]models.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"status": status})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
	)
	return err
}

func (r *taskRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
