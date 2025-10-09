package handlers

import (
	"MovieManager/character"
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type CharacterWithIDHandler struct {
	log *zap.Logger
	db  character.CharacterDataBase
}

func NewCharacterWithIDHandler(log *zap.Logger, db character.CharacterDataBase) *CharacterWithIDHandler {
	return &CharacterWithIDHandler{log: log, db: db}
}

func (*CharacterWithIDHandler) Pattern() string {
	return "/characters/{id}"
}

func (h *CharacterWithIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *CharacterWithIDHandler) getById(w http.ResponseWriter, r *http.Request, id uint64) {
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

func (h *CharacterWithIDHandler) update(w http.ResponseWriter, r *http.Request, id uint64) {
	var c character.Character
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusInternalServerError))
		return
	}

	_, err = h.db.Update(id, c)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNoContent))
}

func (h *CharacterWithIDHandler) delete(w http.ResponseWriter, r *http.Request, id uint64) {
	_, err := h.db.Delete(id)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		h.log.Error("ServeHTTP: Error", zap.Error(err), zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.log.Info("ServeHTTP: Finished", zap.String("URL", h.Pattern()), zap.String("Method", r.Method), zap.Int("Status", http.StatusNoContent))
}
