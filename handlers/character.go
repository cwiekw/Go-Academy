package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type CharacterHandler struct {
	log *zap.Logger
}

func NewCharacterHandler(log *zap.Logger) *CharacterHandler {
	return &CharacterHandler{log: log}
}

func (*CharacterHandler) Pattern() string {
	return "/characters"
}

func (h *CharacterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "All characters on GET")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodPost:
		fmt.Fprintf(w, "Creating new character")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
}

type CharacterWithIDHandler struct {
	log *zap.Logger
}

func NewCharacterWithIDHandler(log *zap.Logger) *CharacterWithIDHandler {
	return &CharacterWithIDHandler{log: log}
}

func (*CharacterWithIDHandler) Pattern() string {
	return "/characters/{id}"
}

func (h *CharacterWithIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ServeHTTP: Started", zap.String("URL", h.Pattern()), zap.String("Method", r.Method))
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Character by ID on GET")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodPut:
		fmt.Fprintf(w, "Updating character by ID")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	case http.MethodDelete:
		fmt.Fprintf(w, "Deleting character by ID")
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusMethodNotAllowed))
		return
	}
}
