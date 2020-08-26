package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mavengopher/gouserview/internal/blog"
	"go.mongodb.org/mongo-driver/mongo"
)

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

//RefLink:https://kb.objectrocket.com/mongo-db/how-to-insert-mongodb-documents-from-json-using-the-golang-driver-457
// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *mongo.Database) error {
	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)

	col := db.Collection("blog")
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("../../internal/schema/blogs.json")
	if err != nil {
		return err
	}

	// Declare an empty slice for the MongoFields docs
	var docs []blog.Blog

	// Unmarshal the encoded JSON byte string into the slice
	err = json.Unmarshal(byteValues, &docs)
	if err != nil {
		return err
	}
	// Iterate the slice of MongoDB struct docs
	for i := range docs {

		// Put the document element in a new variable
		doc := docs[i]
		fmt.Println("ndoc _id:", doc.ID)
		// fmt.Println("doc Field Str:", doc.ID)

		// Call the InsertOne() method and pass the context and doc objects
		result, insertErr := col.InsertOne(ctx, doc)

		// Check for any insertion errors
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
		} else {
			fmt.Println("InsertOne() API result:", result)
		}
	}

	return nil
}
