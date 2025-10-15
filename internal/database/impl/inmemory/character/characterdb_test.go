package character

import (
	"MovieManager/internal/entity/character"
	"testing"

	"github.com/stretchr/testify/assert"
)

const CHARACTER = "Character"
const MOVIE_ID = uint64(12345)
const OTHER_CHARACTER = "Other character"
const OTHER_MOVIE_ID = uint64(98765)

func initDB() *InMemoryCharacterDataBase {
	db := &InMemoryCharacterDataBase{}
	db.db.Store(uint64(1), character.Character{
		Id:      uint64(1),
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	})
	return db
}

func createOtherCharacter() character.Character {
	return character.Character{
		Name:    OTHER_CHARACTER,
		MovieId: OTHER_MOVIE_ID,
	}
}

func TestCharacterDataBase_GetAllCharacters(t *testing.T) {
	characterDb := initDB()

	characters := characterDb.GetAll()

	assert.Len(t, characters, 1, "Should return list of Characters")
}

func TestCharacterDataBase_GetCharacterById_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.GetById(1000)

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_GetCharacterById(t *testing.T) {
	characterDb := initDB()

	c, err := characterDb.GetById(uint64(1))

	expected := character.Character{
		Id:      uint64(1),
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}

	assert.Nil(t, err, "No error")
	assert.Equal(t, expected, c, "Should return Character for ID 1")
}

func TestCharacterDataBase_AddCharacter(t *testing.T) {
	characterDb := initDB()

	newCharacter := characterDb.Add(createOtherCharacter())
	assert.Greater(t, newCharacter.Id, uint64(0), "Should add new character to db and return it. Generated ID should be greater than 0")
	assert.Equal(t, OTHER_CHARACTER, newCharacter.Name, "Should add new character to db and return it. Name should be equal to provided one")
	assert.Equal(t, OTHER_MOVIE_ID, newCharacter.MovieId, "Should add new character to db and return it. Year should be equal to provided one")
	assert.Len(t, characterDb.GetAll(), 2, "Should be 2 characters in DB")
}

func TestCharacterDataBase_UpdateCharacter_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.Update(1000, createOtherCharacter())

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_UpdateCharacter(t *testing.T) {
	characterDb := initDB()

	isCharacterUpdated, err := characterDb.Update(1, createOtherCharacter())

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, isCharacterUpdated, "Should update Character for ID 1 and return true")

	v, _ := characterDb.db.Load(uint64(1))
	updatedCharacter := v.(character.Character)

	assert.Equal(t, updatedCharacter.Id, uint64(1), "ID should stay the same")
	assert.Equal(t, updatedCharacter.Name, OTHER_CHARACTER, "Name should be updated")
	assert.Equal(t, updatedCharacter.MovieId, OTHER_MOVIE_ID, "MovieId should be updated")
}

func TestCharacterDataBase_DeleteCharacter_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.Delete(1000)

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_DeleteCharacter(t *testing.T) {
	characterDb := initDB()

	characterDeleted, err := characterDb.Delete(uint64(1))

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, characterDeleted, "Should deleted Character for ID 1 and return true")
	assert.Len(t, characterDb.GetAll(), 0, "DB should be empty after deleting character")
}
