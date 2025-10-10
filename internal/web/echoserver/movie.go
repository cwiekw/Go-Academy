package echoserver

import (
	"MovieManager/internal/web/api"
	"MovieManager/internal/web/mapper"
	"context"
)

func (s Server) GetMovies(_ context.Context, _ api.GetMoviesRequestObject) (api.GetMoviesResponseObject, error) {
	movies := s.movieDb.GetAll()

	var res api.GetMovies200JSONResponse

	res = make([]api.Movie, len(movies))

	for idx, m := range movies {
		res[idx] = mapper.MapMovieEntityToDto(m)
	}

	return res, nil
}

func (s Server) GetMoviesMovieId(ctx context.Context, request api.GetMoviesMovieIdRequestObject) (api.GetMoviesMovieIdResponseObject, error) {

	if request.MovieId == 1 {
		return api.GetMoviesMovieId404Response{}, nil
	}

	return api.GetMoviesMovieId200JSONResponse{
		Name: "Matrix",
		Year: 1999,
	}, nil
}

func (s Server) PostMovies(ctx context.Context, request api.PostMoviesRequestObject) (api.PostMoviesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutMoviesMovieId(ctx context.Context, request api.PutMoviesMovieIdRequestObject) (api.PutMoviesMovieIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteMoviesMovieId(ctx context.Context, request api.DeleteMoviesMovieIdRequestObject) (api.DeleteMoviesMovieIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
