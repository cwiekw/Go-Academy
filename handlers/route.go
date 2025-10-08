package handlers

import (
	"net/http"

	"go.uber.org/fx"
)

const ROUTES_TAG = `group:"routes"`

type Route interface {
	http.Handler

	Pattern() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(ROUTES_TAG),
	)
}
