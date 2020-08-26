package schema

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed from this slice once they have been ran in
// production.
//
// Including the queries directly in this file has the same pros/cons mentioned
// in seeds.go

var collectionList = []string{"blog"}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *mongo.Database) error {
	//https://www.tutorialspoint.com/mongodb/mongodb_create_collection.htm
	// opt := options.CreateIndexesOptions{}
	db.Collection("blog")
	return nil
}
