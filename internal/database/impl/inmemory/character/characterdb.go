package character

import (
	"MovieManager/internal/database"
	"MovieManager/internal/database/impl"
	"MovieManager/internal/entity/character"
	"sync"
)

type InMemoryCharacterDataBase struct {
	db sync.Map
}

func New() *InMemoryCharacterDataBase {
	return &InMemoryCharacterDataBase{}
}

func (cdb *InMemoryCharacterDataBase) GetAll() []character.Character {
	result := make([]character.Character, 0)

	cdb.db.Range(func(_, value any) bool {
		result = append(result, value.(character.Character))
		return true
	})

	return result
}

func (cdb *InMemoryCharacterDataBase) GetById(id uint64) (character.Character, error) {
	c, exists := cdb.db.Load(id)

	if !exists {
		return character.Character{}, database.NewEntityDoesNotExistError("Character", id)
	}

	return c.(character.Character), nil
}

func (cdb *InMemoryCharacterDataBase) Add(c character.Character) character.Character {
	id := impl.GenerateId()
	c.Id = id
	cdb.db.Store(id, c)
	return c
}

func (cdb *InMemoryCharacterDataBase) Update(id uint64, u character.Character) (bool, error) {
	c, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	c.Name = u.Name
	c.MovieId = u.MovieId

	cdb.db.Store(id, c)

	return true, nil
}

func (cdb *InMemoryCharacterDataBase) Delete(id uint64) (bool, error) {
	_, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	cdb.db.Delete(id)

	return true, nil
}
