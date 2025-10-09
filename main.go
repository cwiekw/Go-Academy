package main

import (
	"MovieManager/character"
	"MovieManager/handlers"
	"MovieManager/movie"
	"MovieManager/server"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewHttpServer,
			fx.Annotate(movie.NewInMemoryMovieDataBase, fx.As(new(movie.MovieDataBase))),
			fx.Annotate(character.NewInMemoryCharacterDataBase, fx.As(new(character.CharacterDataBase))),
			fx.Annotate(server.NewServeMux, fx.ParamTags(handlers.ROUTES_TAG)),
			handlers.AsRoute(handlers.NewHealthCheckHandler),
			handlers.AsRoute(handlers.NewMovieHandler),
			handlers.AsRoute(handlers.NewMMovieWithIDHandler),
			handlers.AsRoute(handlers.NewCharacterHandler),
			handlers.AsRoute(handlers.NewCharacterWithIDHandler),
			zap.NewExample,
		),
		fx.Invoke(func(server *http.Server) {}),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
