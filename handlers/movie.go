package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type MovieHandler struct {
	log *zap.Logger
}

func NewMovieHandler(log *zap.Logger) *MovieHandler {
	return &MovieHandler{log: log}
}

func (*MovieHandler) Pattern() string {
	return "/movies"
}

func (h *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "All movies on GET")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodPost:
		fmt.Fprintf(w, "Creating new movie")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
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
