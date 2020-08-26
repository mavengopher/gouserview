package handlers

import (
	"log"
	"net/http"

	"github.com/mavengopher/gouserview/internal/platform/web"
	"go.mongodb.org/mongo-driver/mongo"
)

// API constructs an http.Handler with all application routes defined.
func API(db *mongo.Database, log *log.Logger) http.Handler {

	app := web.NewApp(log)

	p := Blogs{db: db, log: log}

	app.Handle(http.MethodGet, "/api/v1/blogger/blog/{type}/{page}/{limit}", p.ListByType)
	app.Handle(http.MethodGet, "/api/v1/blogger/blog/{type}/{tag}/{page}/{limit}", p.ListByTag)
	app.Handle(http.MethodGet, "/api/v1/blogger/blog/{id}", p.Retrieve)

	return app
}
