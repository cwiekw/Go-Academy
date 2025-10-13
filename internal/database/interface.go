package database

import (
	"MovieManager/internal/entity/character"
	"MovieManager/internal/entity/movie"
)

type MovieDataBase interface {
	GetAll() []movie.Movie
	GetById(id uint64) (movie.Movie, error)
	Add(m movie.Movie) movie.Movie
	Update(id uint64, u movie.Movie) (bool, error)
	Delete(id uint64) (bool, error)
}

type CharacterDataBase interface {
	GetAll() []character.Character
	GetById(id uint64) (character.Character, error)
	Add(m character.Character) uint64
	Update(id uint64, u character.Character) (bool, error)
	Delete(id uint64) (bool, error)
}
