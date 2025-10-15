package echoserver

import (
	"MovieManager/internal/web/api"
	"MovieManager/internal/web/mapper"
	"context"

	"go.uber.org/zap"
)

func (s Server) GetMovies(_ context.Context, _ api.GetMoviesRequestObject) (api.GetMoviesResponseObject, error) {
	s.log.Info("GetMovies", zap.String("action", "started"))
	movies := s.movieDb.GetAll()

	res := make([]api.Movie, len(movies))

	for idx, m := range movies {
		res[idx] = mapper.MapMovieEntityToDto(m)
	}

	s.log.Info("GetMovies", zap.String("action", "finished"))
	return api.GetMovies200JSONResponse(res), nil
}

func (s Server) GetMoviesMovieId(_ context.Context, request api.GetMoviesMovieIdRequestObject) (api.GetMoviesMovieIdResponseObject, error) {
	s.log.Info("GetMoviesMovieId", zap.String("action", "started"))
	m, err := s.movieDb.GetById(request.MovieId)
	if err != nil {
		s.log.Info("GetMoviesMovieId", zap.String("action", "failed"), zap.Error(err))
		return api.GetMoviesMovieId404Response{}, nil
	}

	s.log.Info("GetMoviesMovieId", zap.String("action", "finished"))
	return api.GetMoviesMovieId200JSONResponse(mapper.MapMovieEntityToDto(m)), nil
}

func (s Server) PostMovies(_ context.Context, request api.PostMoviesRequestObject) (api.PostMoviesResponseObject, error) {
	s.log.Info("PostMovies", zap.String("action", "started"))
	m := mapper.MapMovieDtoToEntity(*request.Body)

	nm := s.movieDb.Add(m)

	s.log.Info("PostMovies", zap.String("action", "finished"))
	return api.PostMovies201JSONResponse(mapper.MapMovieEntityToDto(nm)), nil
}

func (s Server) PutMoviesMovieId(_ context.Context, request api.PutMoviesMovieIdRequestObject) (api.PutMoviesMovieIdResponseObject, error) {
	s.log.Info("PutMoviesMovieId", zap.String("action", "started"))

	_, err := s.movieDb.Update(request.MovieId, mapper.MapMovieDtoToEntity(*request.Body))
	if err != nil {
		s.log.Info("PutMoviesMovieId", zap.String("action", "failed"), zap.Error(err))
		return api.PutMoviesMovieId404Response{}, nil
	}

	s.log.Info("PutMoviesMovieId", zap.String("action", "finished"))
	return api.PutMoviesMovieId204Response{}, nil
}

func (s Server) DeleteMoviesMovieId(_ context.Context, request api.DeleteMoviesMovieIdRequestObject) (api.DeleteMoviesMovieIdResponseObject, error) {
	s.log.Info("DeleteMoviesMovieId", zap.String("action", "started"))

	_, err := s.movieDb.Delete(request.MovieId)
	if err != nil {
		s.log.Info("DeleteMoviesMovieId", zap.String("action", "failed"), zap.Error(err))
		return api.DeleteMoviesMovieId404Response{}, nil
	}

	s.log.Info("DeleteMoviesMovieId", zap.String("action", "finished"))
	return api.DeleteMoviesMovieId204Response{}, nil
}
