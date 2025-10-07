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

func (c CharacterDataBase) String() string {
	var sb strings.Builder

	for id, character := range c.db {
		sb.WriteString(fmt.Sprintf("%d: %s\n", id, character))
	}

	return sb.String()
}

func (c CharacterDataBase) GetAllCharacters() []Character {
	result := make([]Character, len(c.db))

	idx := 0
	for _, character := range c.db {
		result[idx] = character
		idx++
	}

	return result
}

func (c CharacterDataBase) GetCharacterById(id uint64) (Character, error) {
	character, exists := c.db[id]

	if !exists {
		return character, utils.NewEntityDoesNotExistError("Character", id)
	}

	return character, nil
}

func (c CharacterDataBase) AddCharacter(character Character) uint64 {
	id := utils.GenerateId()
	character.Id = id
	c.db[id] = character
	return id
}

func (c CharacterDataBase) UpdateCharacter(id uint64, character Character) (bool, error) {
	_, err := c.GetCharacterById(id)

	if err != nil {
		return false, err
	}

	updatedCharacter := c.db[id]

	updatedCharacter.Name = character.Name
	updatedCharacter.MovieId = character.MovieId

	c.db[id] = updatedCharacter

	return true, nil
}

func (c CharacterDataBase) DeleteCharacter(id uint64) (bool, error) {
	_, err := c.GetCharacterById(id)

	if err != nil {
		return false, err
	}

	delete(c.db, id)

	return true, nil
}
