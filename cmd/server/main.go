package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/chrishadi/movies/internal/pkg/models"
	"github.com/chrishadi/movies/internal/pkg/validators"
)

const (
	addr   = ":8080"
	dbHost = "localhost"
	dbUser = "pguser"
	dbPass = "pgpass"
	dbName = "movies_db"
	dbPort = 5432
)

type handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *handler {
	return &handler{db: db}
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot open database connection: %s", err)
	}

	if err = db.AutoMigrate(&models.Director{}); err != nil {
		log.Fatalf("Failed to migrate the 'directors' table: %s", err)
	}

	if err = db.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatalf("Failed to migrate the 'movies' table: %s", err)
	}

	h := newHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /directors", h.handlePostToDirectors)
	mux.HandleFunc("POST /movies", h.handlePostToMovies)

	log.Print("Listening at " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func (h *handler) handlePostToDirectors(w http.ResponseWriter, r *http.Request) {
	var director models.Director
	if err := jsonapi.UnmarshalPayload(r.Body, &director); err != nil {
		log.Printf("Cannot unmarshal director payload: %s", err)
		respondWithErrorsJSON(w, http.StatusBadRequest, errors.New("invalid director payload"))
		return
	}

	errs := validators.ValidateDirector(&director)
	if len(errs) > 0 {
		respondWithErrorsJSON(w, http.StatusBadRequest, errs...)
		return
	}

	if result := h.db.Create(&director); result.Error != nil {
		log.Printf("Failed to insert director into database: %s", result.Error)
		respondWithErrorsJSON(w, http.StatusInternalServerError, errors.New("oops, something went wrong"))
		return
	}

	respondWithModelJSON(w, http.StatusCreated, &director)
}

func (h *handler) handlePostToMovies(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := jsonapi.UnmarshalPayload(r.Body, &movie); err != nil {
		log.Printf("Cannot unmarshal movie payload: %s", err)
		respondWithErrorsJSON(w, http.StatusBadRequest, errors.New("invalid movie payload"))
		return
	}

	if result := h.db.Create(&movie); result.Error != nil {
		log.Printf("Failed to insert movie into database: %s", result.Error)
		respondWithErrorsJSON(w, http.StatusInternalServerError, errors.New("oops, something went wrong"))
		return
	}

	respondWithModelJSON(w, http.StatusCreated, &movie)
}

func respondWithErrorsJSON(w http.ResponseWriter, code int, errs ...error) {
	status := strconv.Itoa(code)
	errObjs := make([]*jsonapi.ErrorObject, len(errs))

	for i := range errs {
		errObjs[i] = &jsonapi.ErrorObject{
			Detail: errs[i].Error(),
			Status: status,
		}
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(code)

	if err := jsonapi.MarshalErrors(w, errObjs); err != nil {
		log.Printf("Error marshaling errors response: %s", err)
	}
}

func respondWithModelJSON(w http.ResponseWriter, code int, model interface{}) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(code)

	if err := jsonapi.MarshalPayload(w, model); err != nil {
		log.Printf("Error marshaling model: %s", err)
	}
}
