package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// App is the entrypoint into our application and what controls the context of
// each request. Feel free to add any configuration data/logic on this type.
type App struct {
	log *log.Logger
	mux *chi.Mux
}

// NewApp constructs an App to handle a set of routes.
func NewApp(log *log.Logger) *App {
	r := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return &App{
		log: log,
		mux: r,
	}
}

// Handle associates a handler function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h http.HandlerFunc) {

	a.mux.MethodFunc(method, url, h)
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
