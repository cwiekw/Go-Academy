package movie

import (
	"MovieManager/internal/database"
	"MovieManager/internal/database/impl"
	"MovieManager/internal/entity/movie"
	"fmt"
	"strings"
)

type InMemoryMovieDataBase struct {
	db map[uint64]movie.Movie
}

func New() *InMemoryMovieDataBase {
	return &InMemoryMovieDataBase{
		db: make(map[uint64]movie.Movie),
	}
}

func (mdb *InMemoryMovieDataBase) String() string {
	var sb strings.Builder

	for _, m := range mdb.db {
		sb.WriteString(fmt.Sprintf("%s\n", m))
	}

	return sb.String()
}

func (mdb *InMemoryMovieDataBase) GetAll() []movie.Movie {
	result := make([]movie.Movie, len(mdb.db))

	idx := 0
	for _, m := range mdb.db {
		result[idx] = m
		idx++
	}

	return result
}

func (mdb *InMemoryMovieDataBase) GetById(id uint64) (movie.Movie, error) {
	m, exists := mdb.db[id]

	if !exists {
		return m, database.NewEntityDoesNotExistError("Movie", id)
	}

	return m, nil
}

func (mdb *InMemoryMovieDataBase) Add(m movie.Movie) movie.Movie {
	id := impl.GenerateId()
	m.Id = id
	mdb.db[id] = m
	return mdb.db[id]
}

func (mdb *InMemoryMovieDataBase) Update(id uint64, u movie.Movie) (bool, error) {
	m, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	m.Name = u.Name
	m.Year = u.Year

	mdb.db[id] = m

	return true, nil
}

func (mdb *InMemoryMovieDataBase) Delete(id uint64) (bool, error) {
	_, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	delete(mdb.db, id)

	return true, nil
}
