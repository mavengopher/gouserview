package blog

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var blogtype = map[string]string{
	"home":          "",
	"educational":   "शैक्षणिक",
	"political":     "राजकीय",
	"entertainment": "मनोरंजन",
	"social":        "सामाजिक",
}

// ListByType gets all blogs from the database.
func ListByType(db *mongo.Database, typ string, page, limit int) ([]Blog, error) {

	blogs := []Blog{}

	coll := db.Collection("blog")

	opt := options.Find()
	// Sort by `_id` field descending
	opt.SetSort(bson.D{{"_id", -1}})

	opt.SetSkip(int64(page * limit))

	// Limit by 10 documents only
	opt.SetLimit(int64(limit))

	filter := make(map[string]interface{})
	filter["type"] = typ
	filter["status"] = 1

	cursor, err := coll.Find(context.Background(), filter, opt)
	if err != nil {
		return []Blog{}, nil
	}
	cursor.All(context.Background(), &blogs)
	/* 	for i := range blogs {
	   		blogs[i].Type = blogtype[blogs[i].Type]
	   	}
	*/return blogs, nil
}

// ListByTag gets all blogs from the database based on type and tags.
func ListByTag(db *mongo.Database, typ, tag string, page, limit int) ([]Blog, error) {

	blogs := []Blog{}

	coll := db.Collection("blog")

	opt := options.Find()
	// Sort by `_id` field descending
	opt.SetSort(bson.D{{"_id", -1}})

	opt.SetSkip(int64(page * limit))

	// Limit by 10 documents only
	opt.SetLimit(int64(limit))

	filter := make(map[string]interface{})
	filter["tag"] = tag
	filter["type"] = typ
	filter["status"] = 1
	cursor, err := coll.Find(context.Background(), filter, opt)
	if err != nil {
		return []Blog{}, nil
	}
	cursor.All(context.Background(), &blogs)

	/* 	for i := range blogs {
		blogs[i].Type = blogtype[blogs[i].Type]
	} */
	return blogs, nil
}

//Count ...
func Count(db *mongo.Database, typ, tag string) (int64, error) {

	coll := db.Collection("blog")

	filter := make(bson.M)
	filter["status"] = 1
	if typ != "" {
		filter["type"] = typ
	}
	if tag != "" {
		filter["tag"] = tag
	}

	cnt, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return cnt, nil

}

// Retrieve finds the product identified by a given ID.
func Retrieve(db *mongo.Database, id string) (*Blog, error) {

	coll := db.Collection("blog")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog Blog
	if err := coll.FindOne(context.Background(), bson.D{{"_id", objID}}).Decode(&blog); err != nil {
		return nil, err
	}
	// blog.Type = blogtype[blog.Type]

	return &blog, nil
}
