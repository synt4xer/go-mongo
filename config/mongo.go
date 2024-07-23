package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProvideClient(cfg *Config) (*mongo.Client, error) {
	mongoURI := cfg.MongoDB.URI

	// * Create a MongoDB client configuration
	opts := options.Client().ApplyURI(mongoURI)

	ctx := context.Background()

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	// * should not put the defer disconnect function here
	// defer func(ctx context.Context) {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }(ctx)

	// * ping the client
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	return client, nil
}
