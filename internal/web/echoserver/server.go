package echoserver

import (
	"MovieManager/internal/database"
	"MovieManager/internal/validator"
	"MovieManager/internal/web/api"
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	log              *zap.Logger
	movieDb          database.MovieDataBase
	characterDb      database.CharacterDataBase
	validatorManager validator.CharacterValidatorManager
}

var _ api.StrictServerInterface = (*Server)(nil)

func New(log *zap.Logger, movieDb database.MovieDataBase, characterDb database.CharacterDataBase, validatorManager validator.CharacterValidatorManager) *Server {
	return &Server{
		log:              log,
		movieDb:          movieDb,
		characterDb:      characterDb,
		validatorManager: validatorManager,
	}
}

func NewEchoServer(lc fx.Lifecycle, log *zap.Logger, movieDb database.MovieDataBase, characterDb database.CharacterDataBase, validatorManager *validator.CharacterValidatorManager) *echo.Echo {
	e := echo.New()
	handlers := New(log, movieDb, characterDb, *validatorManager)

	api.RegisterHandlersWithBaseURL(e, api.NewStrictHandler(handlers, []api.StrictMiddlewareFunc{}), "/api/v1")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting HTTP server", zap.String("addr", ":7734"))
			go func() {
				err := e.Start(":7734")
				if err != nil {
					e.Logger.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping HTTP server")
			err := e.Shutdown(ctx)
			if err != nil {
				return err
			}

			return nil
		},
	})

	return e
}
