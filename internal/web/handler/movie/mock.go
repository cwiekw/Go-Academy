package movie

import (
	"MovieManager/internal/database"
	emovie "MovieManager/internal/entity/movie"

	"github.com/stretchr/testify/mock"
)

type mockMovieDataBase struct {
	mock.Mock
}

func newMockMovieDataBase() *mockMovieDataBase {
	m := &mockMovieDataBase{}

	id := uint64(1)
	wrongId := uint64(999)
	movie := emovie.Movie{
		Name: "Movie1",
		Year: 2001,
	}
	movieWithId := emovie.Movie{
		Id:   id,
		Name: "Movie1",
		Year: 2001,
	}

	m.On("GetAll").Return([]emovie.Movie{
		movieWithId,
		{
			Id:   uint64(2),
			Name: "Movie2",
			Year: 2002,
		},
	})

	m.On("Add", movie).Return(uint64(1))

	m.On("GetById", id).Return(movieWithId, nil)
	m.On("GetById", wrongId).Return(emovie.Movie{}, database.NewEntityDoesNotExistError("Movie", wrongId))

	m.On("Update", id, movie).Return(true, nil)
	m.On("Update", wrongId, movie).Return(false, database.NewEntityDoesNotExistError("Movie", wrongId))

	m.On("Delete", id).Return(true, nil)
	m.On("Delete", wrongId).Return(false, database.NewEntityDoesNotExistError("Movie", wrongId))

	return m
}

func (m *mockMovieDataBase) GetAll() []emovie.Movie {
	args := m.Called()
	return args.Get(0).([]emovie.Movie)
}

func (m *mockMovieDataBase) GetById(id uint64) (emovie.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(emovie.Movie), args.Error(1)
}

func (m *mockMovieDataBase) Add(em emovie.Movie) uint64 {
	args := m.Called(em)
	return args.Get(0).(uint64)
}

func (m *mockMovieDataBase) Update(id uint64, u emovie.Movie) (bool, error) {
	args := m.Called(id, u)
	return args.Bool(0), args.Error(1)
}

func (m *mockMovieDataBase) Delete(id uint64) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
