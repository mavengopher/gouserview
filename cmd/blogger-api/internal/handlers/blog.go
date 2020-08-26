package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mavengopher/gouserview/internal/blog"
	"go.mongodb.org/mongo-driver/mongo"
)

// Blogs defines all of the handlers related to blogs. It holds the
// application state needed by the handler methods.
type Blogs struct {
	db  *mongo.Database
	log *log.Logger
}

// ListByType gets all Blogs from the service layer and encodes them for the
// client response.
func (p *Blogs) ListByType(w http.ResponseWriter, r *http.Request) {

	typ := chi.URLParam(r, "type")
	page, _ := strconv.Atoi(chi.URLParam(r, "page"))
	limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))
	if limit <= 0 {
		limit = 6
	}

	list, err := blog.ListByType(p.db, typ, page, limit)
	if err != nil {
		p.log.Println("listing Blogs", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cnt, err := blog.Count(p.db, typ, "")
	if err != nil {
		p.log.Println("Count query Error :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := make(map[string]interface{})
	b["blogs"] = list
	b["count"] = cnt

	data, err := json.Marshal(b)
	if err != nil {
		p.log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.log.Println("error writing result", err)
	}
}

// ListByTag gets all Blogs from the service layer and encodes them for the
// client response.
func (p *Blogs) ListByTag(w http.ResponseWriter, r *http.Request) {

	typ := chi.URLParam(r, "type")
	tag := chi.URLParam(r, "tag")
	page, _ := strconv.Atoi(chi.URLParam(r, "page"))
	limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))
	if limit <= 0 {
		limit = 6
	}

	list, err := blog.ListByTag(p.db, typ, tag, page, limit)
	if err != nil {
		p.log.Println("listing Blogs", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cnt, err := blog.Count(p.db, typ, "")
	if err != nil {
		p.log.Println("Count query Error :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := make(map[string]interface{})
	b["blogs"] = list
	b["count"] = cnt

	data, err := json.Marshal(b)
	if err != nil {
		p.log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.log.Println("error writing result", err)
	}
}

// Retrieve finds a single product identified by an ID in the request URL.
func (p *Blogs) Retrieve(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	prod, err := blog.Retrieve(p.db, id)
	if err != nil {
		p.log.Println("getting product", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(prod)
	if err != nil {
		p.log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.log.Println("error writing result", err)
	}
}
