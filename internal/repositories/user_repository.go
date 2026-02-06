package repositories

import (
	"context"

	"Concurrent_Task_Management_System/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindAll(ctx context.Context) ([]models.User, error)
	UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return nil
}


func (r *userRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUserID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
	)
	return err
}

func (r *userRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
