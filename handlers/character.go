package handlers

import (
	"MovieManager/character"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type CharacterHandler struct {
	log *zap.Logger
	db  character.CharacterDataBase
}

func NewCharacterHandler(log *zap.Logger, db character.CharacterDataBase) *CharacterHandler {
	return &CharacterHandler{log: log, db: db}
}

func (*CharacterHandler) Pattern() string {
	return "/characters"
}

func (h *CharacterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *CharacterHandler) getAll(w http.ResponseWriter, r *http.Request) {
	characters := h.db.GetAll()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(characters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusOK))
}

func (h *CharacterHandler) add(w http.ResponseWriter, r *http.Request) {
	var c character.Character
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	id := h.db.Add(c)

	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, "%d", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusCreated))
}
