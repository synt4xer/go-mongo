package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/synt4xer/go-mongo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

// * constructor
func NewUserRepository(repo *MongoRepository) (*UserRepository, error) {
	col, err := repo.Collection("users")
	if err != nil {
		return nil, err
	}

	return &UserRepository{collection: col}, nil
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	now := time.Now().UTC()

	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = &now
	user.UpdatedAt = &now

	data, err := bson.Marshal(user)

	if err != nil {
		return nil, err
	}

	_, err = r.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	now := time.Now().UTC()

	user.UpdatedAt = &now

	// no need to be marshal because mongo update gonna marshal it
	// data, err := bson.Marshal(user)

	// if err != nil {
	// 	return nil, err
	// }

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": user}

	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
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

func (r *UserRepository) GetAll(ctx context.Context, search string) ([]domain.User, error) {
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

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &user, nil
}
