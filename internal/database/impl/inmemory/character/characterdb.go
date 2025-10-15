package character

import (
	"MovieManager/internal/database"
	"MovieManager/internal/database/impl"
	"MovieManager/internal/entity/character"
	"fmt"
	"strings"
)

type InMemoryCharacterDataBase struct {
	db map[uint64]character.Character
}

func New() InMemoryCharacterDataBase {
	return InMemoryCharacterDataBase{
		db: make(map[uint64]character.Character),
	}
}

func (cdb InMemoryCharacterDataBase) String() string {
	var sb strings.Builder

	for id, c := range cdb.db {
		sb.WriteString(fmt.Sprintf("%d: %s\n", id, c))
	}

	return sb.String()
}

func (cdb InMemoryCharacterDataBase) GetAll() []character.Character {
	result := make([]character.Character, len(cdb.db))

	idx := 0
	for _, c := range cdb.db {
		result[idx] = c
		idx++
	}

	return result
}

func (cdb InMemoryCharacterDataBase) GetById(id uint64) (character.Character, error) {
	c, exists := cdb.db[id]

	if !exists {
		return c, database.NewEntityDoesNotExistError("Character", id)
	}

	return c, nil
}

func (cdb InMemoryCharacterDataBase) Add(c character.Character) character.Character {
	id := impl.GenerateId()
	c.Id = id
	cdb.db[id] = c
	return cdb.db[id]
}

func (cdb InMemoryCharacterDataBase) Update(id uint64, u character.Character) (bool, error) {
	c, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	c.Name = u.Name
	c.MovieId = u.MovieId

	cdb.db[id] = c

	return true, nil
}

func (cdb InMemoryCharacterDataBase) Delete(id uint64) (bool, error) {
	_, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	delete(cdb.db, id)

	return true, nil
}
