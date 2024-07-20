package repository

import (
	"context"
	"fmt"

	"github.com/synt4xer/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

// * constructor
func NewMongoRepository(ctx context.Context, client *mongo.Client, config *config.Config) (*mongoRepository, error) {
	dbName := config.MongoDB.Database

	if dbName == "" {
		return nil, fmt.Errorf("missing database in configuration")
	}

	db := client.Database(dbName)
	return &mongoRepository{client: client, db: db}, nil
}

func (r *mongoRepository) Collection(collectionName string) (*mongo.Collection, error) {
	return r.db.Collection(collectionName), nil
}
