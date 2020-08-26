package database

import (
	"context"
	"time"

	_ "github.com/lib/pq" // The database driver in use.
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config is the required properties to use the database.
type Config struct {
	URI  string
	Name string
}

// Open knows how to open a database connection based on the configuration.
func Open(cfg Config) (*mongo.Database, context.Context, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, ctx, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, ctx, err
	}
	return client.Database(cfg.Name), ctx, err
}
