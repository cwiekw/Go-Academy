package character

import (
	"MovieManager/utils"
	"fmt"
	"strings"
)

type CharacterDataBase struct {
	db map[uint64]Character
}

func NewCharacterDataBase() CharacterDataBase {
	return CharacterDataBase{
		db: make(map[uint64]Character),
	}
}

func (cdb CharacterDataBase) String() string {
	var sb strings.Builder

	for id, c := range cdb.db {
		sb.WriteString(fmt.Sprintf("%d: %s\n", id, c))
	}

	return sb.String()
}

func (cdb CharacterDataBase) GetAll() []Character {
	result := make([]Character, len(cdb.db))

	idx := 0
	for _, c := range cdb.db {
		result[idx] = c
		idx++
	}

	return result
}

func (cdb CharacterDataBase) GetById(id uint64) (Character, error) {
	c, exists := cdb.db[id]

	if !exists {
		return c, utils.NewEntityDoesNotExistError("Character", id)
	}

	return c, nil
}

func (cdb CharacterDataBase) Add(c Character) uint64 {
	id := utils.GenerateId()
	c.Id = id
	cdb.db[id] = c
	return id
}

func (cdb CharacterDataBase) Update(id uint64, u Character) (bool, error) {
	c, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	c.Name = u.Name
	c.MovieId = u.MovieId

	cdb.db[id] = c

	return true, nil
}

func (cdb CharacterDataBase) Delete(id uint64) (bool, error) {
	_, err := cdb.GetById(id)

	if err != nil {
		return false, err
	}

	delete(cdb.db, id)

	return true, nil
}
