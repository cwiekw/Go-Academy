package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type HealthCheckHandler struct {
	log *zap.Logger
}

func NewHealthCheckHandler(log *zap.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{log: log}
}

func (*HealthCheckHandler) Pattern() string {
	return "/healthcheck"
}

func (*HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}
