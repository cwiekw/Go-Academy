package handlers

import (
	"MovieManager/movie"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type MovieHandler struct {
	log *zap.Logger
	db  movie.MovieDataBase
}

func NewMovieHandler(log *zap.Logger, db movie.MovieDataBase) *MovieHandler {
	return &MovieHandler{log: log, db: db}
}

func (*MovieHandler) Pattern() string {
	return "/movies"
}

func (h *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))
	switch r.Method {
	case http.MethodGet:
		h.getAll(w, r)
		return
	case http.MethodPost:
		h.add(w, r)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
}

func (h *MovieHandler) getAll(w http.ResponseWriter, r *http.Request) {
	movies := h.db.GetAll()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
}

func (h *MovieHandler) add(w http.ResponseWriter, r *http.Request) {
	var m movie.Movie
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	id := h.db.Add(m)

	_, err = fmt.Fprintf(w, "%d", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusCreated))
}

type MovieWithIDHandler struct {
	log *zap.Logger
}

func NewMMovieWithIDHandler(log *zap.Logger) *MovieWithIDHandler {
	return &MovieWithIDHandler{log: log}
}

func (*MovieWithIDHandler) Pattern() string {
	return "/movies/{id}"
}

func (h *MovieWithIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Movie by ID on GET")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodPut:
		fmt.Fprintf(w, "Updating movie by ID")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodDelete:
		fmt.Fprintf(w, "Deleting movie by ID")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
}
