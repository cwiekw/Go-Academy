package mapper

import (
	"MovieManager/internal/entity/movie"
	"MovieManager/internal/web/api"
)

func MapMovieDtoToEntity(m api.Movie) movie.Movie {
	var id uint64
	if m.Id == nil {
		id = 0
	} else {
		id = *m.Id
	}

	return movie.Movie{
		Id:   id,
		Name: m.Name,
		Year: m.Year,
	}
}

func MapMovieEntityToDto(m movie.Movie) api.Movie {
	var id *uint64
	if m.Id == 0 {
		id = nil
	} else {
		id = &m.Id
	}

	return api.Movie{
		Id:   id,
		Name: m.Name,
		Year: m.Year,
	}
}
