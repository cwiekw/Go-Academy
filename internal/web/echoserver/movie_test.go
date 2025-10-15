package echoserver

import (
	"MovieManager/internal/web/api"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_GetMovies(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetMovies(c, api.GetMoviesRequestObject{})

	expected := api.GetMovies200JSONResponse([]api.Movie{
		{
			Id:   &ID,
			Name: "Movie1",
			Year: 2001,
		},
		{
			Id:   &WRONG_ID,
			Name: "Movie2",
			Year: 2002,
		},
	})

	assert.Equal(t, expected, res, "Should return all movies")
}

func TestServer_GetMoviesMovieId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetMoviesMovieId(c, api.GetMoviesMovieIdRequestObject{MovieId: ID})

	expected := api.GetMoviesMovieId200JSONResponse(api.Movie{
		Id:   &ID,
		Name: "Movie1",
		Year: 2001,
	})

	assert.Equal(t, expected, res, "Should return movie by ID")
}

func TestServer_GetMoviesMovieId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetMoviesMovieId(c, api.GetMoviesMovieIdRequestObject{MovieId: WRONG_ID})

	expected := api.GetMoviesMovieId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}

func TestServer_PostMovies(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PostMovies(c, api.PostMoviesRequestObject{Body: &api.Movie{Name: "Movie1", Year: 2001}})

	expected := api.PostMovies201JSONResponse(api.Movie{
		Id:   &ID,
		Name: "Movie1",
		Year: 2001,
	})

	assert.Equal(t, expected, res, "Should add new movie and return it with 201 entity")
}

func TestServer_PutMoviesMovieId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PutMoviesMovieId(c, api.PutMoviesMovieIdRequestObject{
		MovieId: ID,
		Body: &api.Movie{
			Name: "Movie1",
			Year: 2001,
		}})

	expected := api.PutMoviesMovieId204Response{}

	assert.Equal(t, expected, res, "Should update movie and return 204 entity")
}

func TestServer_PutMoviesMovieId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PutMoviesMovieId(c, api.PutMoviesMovieIdRequestObject{
		MovieId: WRONG_ID,
		Body: &api.Movie{
			Name: "Movie1",
			Year: 2001,
		}})

	expected := api.PutMoviesMovieId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}

func TestServer_DeleteMoviesMovieId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.DeleteMoviesMovieId(c, api.DeleteMoviesMovieIdRequestObject{MovieId: ID})

	expected := api.DeleteMoviesMovieId204Response{}

	assert.Equal(t, expected, res, "Should delete movie and return 204 entity")
}

func TestServer_DeleteMoviesMovieId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.DeleteMoviesMovieId(c, api.DeleteMoviesMovieIdRequestObject{MovieId: WRONG_ID})

	expected := api.DeleteMoviesMovieId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}
