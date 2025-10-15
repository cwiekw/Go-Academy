package movie

import (
	"MovieManager/internal/database"
	"MovieManager/internal/database/impl"
	"MovieManager/internal/entity/movie"
	"sync"
)

type InMemoryMovieDataBase struct {
	db sync.Map
}

func New() *InMemoryMovieDataBase {
	return &InMemoryMovieDataBase{}
}

func (mdb *InMemoryMovieDataBase) GetAll() []movie.Movie {
	result := make([]movie.Movie, 0)

	mdb.db.Range(func(_, value any) bool {
		result = append(result, value.(movie.Movie))
		return true
	})

	return result
}

func (mdb *InMemoryMovieDataBase) GetById(id uint64) (movie.Movie, error) {
	m, exists := mdb.db.Load(id)

	if !exists {
		return movie.Movie{}, database.NewEntityDoesNotExistError("Movie", id)
	}

	return m.(movie.Movie), nil
}

func (mdb *InMemoryMovieDataBase) Add(m movie.Movie) movie.Movie {
	id := impl.GenerateId()
	m.Id = id
	mdb.db.Store(id, m)
	return m
}

func (mdb *InMemoryMovieDataBase) Update(id uint64, u movie.Movie) (bool, error) {
	m, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	m.Name = u.Name
	m.Year = u.Year

	mdb.db.Store(id, m)

	return true, nil
}

func (mdb *InMemoryMovieDataBase) Delete(id uint64) (bool, error) {
	_, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	mdb.db.Delete(id)

	return true, nil
}
