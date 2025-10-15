package mapper

import (
	"MovieManager/internal/entity/character"
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

func MapCharacterDtoToEntity(c api.Character) character.Character {
	var id uint64
	if c.Id == nil {
		id = 0
	} else {
		id = *c.Id
	}

	return character.Character{
		Id:      id,
		Name:    c.Name,
		MovieId: c.MovieId,
	}
}

func MapCharacterEntityToDto(c character.Character) api.Character {
	var id *uint64
	if c.Id == 0 {
		id = nil
	} else {
		id = &c.Id
	}

	return api.Character{
		Id:      id,
		Name:    c.Name,
		MovieId: c.MovieId,
	}

}
