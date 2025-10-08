package main

import (
	"MovieManager/character"
	"MovieManager/handlers"
	"MovieManager/movie"
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func oldMain() {
	movieDb := movie.NewInMemoryMovieDataBase()
	characterDb := character.NewCharacterDataBase()

	m1 := movieDb.Add(movie.New(
		movie.WithName("Matrix"),
		movie.WithYear(1999),
	))

	characterDb.Add(character.New(
		character.WithName("Neo"),
		character.WithMovieId(m1),
	))

	characterDb.Add(character.New(
		character.WithName("Agent Smith"),
		character.WithMovieId(m1),
	))

	m2 := movieDb.Add(movie.New(
		movie.WithName("Casino Rolaye"),
		movie.WithYear(2002),
	))

	c1 := characterDb.Add(character.New(
		character.WithName("Bames Jond"),
		character.WithMovieId(m2),
	))

	c2 := characterDb.Add(character.New(
		character.WithName("N"),
		character.WithMovieId(m2),
	))

	m3 := movieDb.Add(movie.New(
		movie.WithName("Mandalorian"),
		movie.WithYear(2019),
	))

	c3 := characterDb.Add(character.New(
		character.WithName("Mando"),
		character.WithMovieId(m3),
	))

	_, _ = movieDb.Update(m2, movie.New(
		movie.WithName("Casino Royale"),
		movie.WithYear(2006),
	))

	_, _ = characterDb.Update(c1, character.New(
		character.WithName("James Bond"),
		character.WithMovieId(m2),
	))

	_, _ = characterDb.Update(c2, character.New(
		character.WithName("M"),
		character.WithMovieId(m2),
	))

	_, _ = movieDb.Delete(m3)
	_, _ = characterDb.Delete(c3)

	allMovies := movieDb.GetAll()
	fmt.Println("Movies:")
	for _, movieEntry := range allMovies {
		fmt.Println(movieEntry)
	}

	fmt.Println()
	fmt.Println()

	allCharacters := characterDb.GetAll()
	fmt.Println("Characters:")
	for _, characterEntry := range allCharacters {
		movieForCharacter, _ := movieDb.GetById(characterEntry.MovieId)
		fmt.Printf("%s: %s\n", characterEntry.Name, movieForCharacter.Name)
	}
}

func NewServeMux(routes []handlers.Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":7734", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func main() {
	fx.New(
		fx.Provide(
			NewHttpServer,
			fx.Annotate(
				movie.NewInMemoryMovieDataBase,
				fx.As(new(movie.MovieDataBase)),
			),
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(handlers.ROUTES_TAG),
			),
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
