package repository

import (
	"fmt"

	"github.com/synt4xer/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

// * constructor
func NewMongoRepository(client *mongo.Client, config *config.Config) (*MongoRepository, error) {
	dbName := config.MongoDB.Database

	if dbName == "" {
		return nil, fmt.Errorf("missing database in configuration")
	}

	db := client.Database(dbName)
	return &MongoRepository{client: client, db: db}, nil
}

func (r *MongoRepository) Collection(collectionName string) (*mongo.Collection, error) {
	return r.db.Collection(collectionName), nil
}
