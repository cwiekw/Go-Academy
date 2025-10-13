package character

import (
	"MovieManager/internal/database"
	echaracter "MovieManager/internal/entity/character"

	"github.com/stretchr/testify/mock"
)

type mockCharacterDataBase struct {
	mock.Mock
}

func newMockCharacterDataBase() *mockCharacterDataBase {
	m := &mockCharacterDataBase{}

	id := uint64(1)
	wrongId := uint64(999)
	character := echaracter.Character{
		Name:    "Character1",
		MovieId: uint64(1001),
	}
	characterWithId := echaracter.Character{
		Id:      id,
		Name:    "Character1",
		MovieId: uint64(1001),
	}

	m.On("GetAll").Return([]echaracter.Character{
		characterWithId,
		{
			Id:      uint64(2),
			Name:    "Character2",
			MovieId: uint64(1002),
		},
	})

	m.On("Add", character).Return(characterWithId)

	m.On("GetById", id).Return(characterWithId, nil)
	m.On("GetById", wrongId).Return(echaracter.Character{}, database.NewEntityDoesNotExistError("Character", wrongId))

	m.On("Update", id, character).Return(true, nil)
	m.On("Update", wrongId, character).Return(false, database.NewEntityDoesNotExistError("Character", wrongId))

	m.On("Delete", id).Return(true, nil)
	m.On("Delete", wrongId).Return(false, database.NewEntityDoesNotExistError("Character", wrongId))

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
