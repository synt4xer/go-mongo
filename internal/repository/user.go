package repository

import (
	"context"
	"fmt"

	"github.com/synt4xer/go-mongo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

// * constructor
func NewUserRepository(repo mongoRepository) (*userRepository, error) {
	col, err := repo.Collection("users")
	if err != nil {
		return nil, err
	}

	return &userRepository{collection: col}, nil
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	user.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) Update(ctx context.Context, id string, updates interface{}) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": updates}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}

func (r *userRepository) GetAll(ctx context.Context, search string) ([]domain.User, error) {
	filter := bson.M{}

	if search != "" {
		filter = bson.M{
			"$text": bson.M{
				"$search": search,
			},
		}
	}

	cursor, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []domain.User

	for {

		isExist := cursor.Next(ctx)
		// if err == mongo.ErrNoDocuments {
		// 	break
		// }

		// if err != nil {
		// 	return nil, err
		// }

		if !isExist {
			break
		}

		var user domain.User
		err = cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &user, nil
}
