package blog

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog is an item we publish.
type Blog struct {
	ID          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EventDate   string              `json:"eventdate" bson:"eventdate"`
	Image       string              `json:"image" bson:"image"`
	Headline    string              `json:"headline" bson:"headline"`
	Information string              `json:"information" bson:"information"`
	Location    string              `json:"location" bson:"location"`
	Type        string              `json:"type" bson:"type"`
	Tag         string              `json:"tag" bson:"tag"`
	Status      int                 `json:"status" bson:"status"`
	CreatedAt   time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt" bson:"updatedAt"`
}
