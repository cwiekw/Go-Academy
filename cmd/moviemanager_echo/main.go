package main

import (
	"MovieManager/internal/database"
	dbcharacter "MovieManager/internal/database/impl/inmemory/character"
	dbmovie "MovieManager/internal/database/impl/inmemory/movie"
	"MovieManager/internal/initialization"
	"MovieManager/internal/validator"
	"MovieManager/internal/web/echoserver"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

//go:generate go tool oapi-codegen -config ../../api/codegen-server.yml ../../api/openapi.yml

func main() {
	fx.New(
		fx.Provide(
			echoserver.NewEchoServer,
			fx.Annotate(dbmovie.New, fx.As(new(database.MovieDataBase))),
			fx.Annotate(dbcharacter.New, fx.As(new(database.CharacterDataBase))),
			validator.NewCharacterValidatorManager,
			zap.NewExample,
		),
		fx.Invoke(
			func(server *echo.Echo) {},
			initialization.Init,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
