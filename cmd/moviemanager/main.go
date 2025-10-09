package main

import (
	"MovieManager/internal/database"
	dbcharacter "MovieManager/internal/database/impl/inmemory/character"
	dbmovie "MovieManager/internal/database/impl/inmemory/movie"
	"MovieManager/internal/web/handler"
	hcharacter "MovieManager/internal/web/handler/character"
	hmovie "MovieManager/internal/web/handler/movie"
	"MovieManager/internal/web/healthcheck"
	"MovieManager/internal/web/server"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewHttpServer,
			fx.Annotate(dbmovie.New, fx.As(new(database.MovieDataBase))),
			fx.Annotate(dbcharacter.New, fx.As(new(database.CharacterDataBase))),
			fx.Annotate(server.NewServeMux, fx.ParamTags(handler.ROUTES_TAG)),
			handler.AsRoute(healthcheck.NewHealthCheckHandler),
			handler.AsRoute(hmovie.NewMovieHandler),
			handler.AsRoute(hmovie.NewMMovieWithIDHandler),
			handler.AsRoute(hcharacter.NewCharacterHandler),
			handler.AsRoute(hcharacter.NewCharacterWithIDHandler),
			zap.NewExample,
		),
		fx.Invoke(func(server *http.Server) {}),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
