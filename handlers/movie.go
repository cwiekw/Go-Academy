package handlers

import (
	"MovieManager/movie"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	w.WriteHeader(http.StatusCreated)
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
	db  movie.MovieDataBase
}

func NewMMovieWithIDHandler(log *zap.Logger, db movie.MovieDataBase) *MovieWithIDHandler {
	return &MovieWithIDHandler{log: log, db: db}
}

func (*MovieWithIDHandler) Pattern() string {
	return "/movies/{id}"
}

func (h *MovieWithIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, "ID in wrong format", http.StatusBadRequest)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusBadRequest))
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getById(w, r, id)
		return
	case http.MethodPut:
		h.update(w, r, id)
		return
	case http.MethodDelete:
		h.delete(w, r, id)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
}

func (h *MovieWithIDHandler) getById(w http.ResponseWriter, r *http.Request, id uint64) {
	res, err := h.db.GetById(id)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNotFound))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
}

func (h *MovieWithIDHandler) update(w http.ResponseWriter, r *http.Request, id uint64) {
	var m movie.Movie
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	_, err = h.db.Update(id, m)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNoContent))
}

func (h *MovieWithIDHandler) delete(w http.ResponseWriter, r *http.Request, id uint64) {
	_, err := h.db.Delete(id)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNoContent))
}
