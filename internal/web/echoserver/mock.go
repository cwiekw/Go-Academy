package echoserver

import (
	"MovieManager/internal/database"
	echaracter "MovieManager/internal/entity/character"
	emovie "MovieManager/internal/entity/movie"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

var ID = uint64(1)
var WRONG_ID = uint64(999)

func newServer() *Server {
	return New(zap.NewNop(), newMockMovieDataBase(), newMockCharacterDataBase())
}

type mockMovieDataBase struct {
	mock.Mock
}

func newMockMovieDataBase() *mockMovieDataBase {
	m := &mockMovieDataBase{}

	movie := emovie.Movie{
		Name: "Movie1",
		Year: 2001,
	}
	movieWithId := emovie.Movie{
		Id:   ID,
		Name: "Movie1",
		Year: 2001,
	}

	m.On("GetAll").Return([]emovie.Movie{
		movieWithId,
		{
			Id:   WRONG_ID,
			Name: "Movie2",
			Year: 2002,
		},
	})

	m.On("Add", movie).Return(movieWithId)

	m.On("GetById", ID).Return(movieWithId, nil)
	m.On("GetById", WRONG_ID).Return(emovie.Movie{}, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

	m.On("Update", ID, movie).Return(true, nil)
	m.On("Update", WRONG_ID, movie).Return(false, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

	m.On("Delete", ID).Return(true, nil)
	m.On("Delete", WRONG_ID).Return(false, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

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

func (m *mockMovieDataBase) Add(em emovie.Movie) emovie.Movie {
	args := m.Called(em)
	return args.Get(0).(emovie.Movie)
}

func (m *mockMovieDataBase) Update(id uint64, u emovie.Movie) (bool, error) {
	args := m.Called(id, u)
	return args.Bool(0), args.Error(1)
}

func (m *mockMovieDataBase) Delete(id uint64) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

type mockCharacterDataBase struct {
	mock.Mock
}

func newMockCharacterDataBase() *mockCharacterDataBase {
	m := &mockCharacterDataBase{}

	character := echaracter.Character{
		Name:    "Character1",
		MovieId: uint64(1001),
	}
	characterWithId := echaracter.Character{
		Id:      ID,
		Name:    "Character1",
		MovieId: uint64(1001),
	}

	m.On("GetAll").Return([]echaracter.Character{
		characterWithId,
		{
			Id:      WRONG_ID,
			Name:    "Character2",
			MovieId: uint64(1002),
		},
	})

	m.On("Add", character).Return(characterWithId)

	m.On("GetById", ID).Return(characterWithId, nil)
	m.On("GetById", WRONG_ID).Return(echaracter.Character{}, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	m.On("Update", ID, character).Return(true, nil)
	m.On("Update", WRONG_ID, character).Return(false, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	m.On("Delete", ID).Return(true, nil)
	m.On("Delete", WRONG_ID).Return(false, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	return m
}

func (m *mockCharacterDataBase) GetAll() []echaracter.Character {
	args := m.Called()
	return args.Get(0).([]echaracter.Character)
}

func (m *mockCharacterDataBase) GetById(id uint64) (echaracter.Character, error) {
	args := m.Called(id)
	return args.Get(0).(echaracter.Character), args.Error(1)
}

func (m *mockCharacterDataBase) Add(em echaracter.Character) echaracter.Character {
	args := m.Called(em)
	return args.Get(0).(echaracter.Character)
}

func (m *mockCharacterDataBase) Update(id uint64, u echaracter.Character) (bool, error) {
	args := m.Called(id, u)
	return args.Bool(0), args.Error(1)
}

func (m *mockCharacterDataBase) Delete(id uint64) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
