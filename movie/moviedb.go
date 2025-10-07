package movie

import (
	"MovieManager/utils"
	"fmt"
	"strings"
)

type MovieDataBase struct {
	db map[uint64]Movie
}

func NewMovieDataBase() MovieDataBase {
	return MovieDataBase{
		db: make(map[uint64]Movie),
	}
}

func (m MovieDataBase) String() string {
	var sb strings.Builder

	for _, movie := range m.db {
		sb.WriteString(fmt.Sprintf("%s\n", movie))
	}

	return sb.String()
}

func (m MovieDataBase) GetAllMovies() ([]Movie, error) {
	if m.db == nil {
		return nil, utils.NewDBNotInitializedError("Movie")
	}

	result := make([]Movie, len(m.db))

	idx := 0
	for _, movie := range m.db {
		result[idx] = movie
		idx++
	}

	return result, nil
}

func (m MovieDataBase) GetMovieById(id uint64) (Movie, error) {
	if m.db == nil {
		return Movie{}, utils.NewDBNotInitializedError("Movie")
	}

	movie, exists := m.db[id]

	if !exists {
		return movie, utils.NewEntityDoesNotExistError("Movie", id)
	}

	return movie, nil
}

func (m MovieDataBase) AddMovie(movie Movie) (uint64, error) {
	if m.db == nil {
		return 0, utils.NewDBNotInitializedError("Movie")
	}

	id := utils.GenerateId()
	movie.Id = id
	m.db[id] = movie
	return id, nil
}

func (m MovieDataBase) UpdateMovie(id uint64, movie Movie) (bool, error) {
	if m.db == nil {
		return false, utils.NewDBNotInitializedError("Movie")
	}

	_, err := m.GetMovieById(id)

	if err != nil {
		return false, err
	}

	updatedMovie := m.db[id]

	updatedMovie.Name = movie.Name
	updatedMovie.Year = movie.Year

	m.db[id] = updatedMovie

	return true, nil
}

func (m MovieDataBase) DeleteMovie(id uint64) (bool, error) {
	if m.db == nil {
		return false, utils.NewDBNotInitializedError("Movie")
	}

	_, err := m.GetMovieById(id)

	if err != nil {
		return false, err
	}

	delete(m.db, id)

	return true, nil
}
