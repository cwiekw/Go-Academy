package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const OTHER_CHARACTER = "Other character"
const OTHER_MOVIE_ID = uint64(98765)

func initDB() CharacterDataBase {
	return CharacterDataBase{
		db: map[uint64]Character{
			1: {
				Id:      1,
				Name:    CHARACTER,
				MovieId: MOVIE_ID,
			},
		},
	}
}

func createOtherCharacter() Character {
	return Character{
		Name:    OTHER_CHARACTER,
		MovieId: OTHER_MOVIE_ID,
	}
}

func TestNewCharacterDataBase(t *testing.T) {
	result := NewCharacterDataBase()

	expected := make(map[uint64]Character)

	assert.Equal(t, expected, result.db, "NewCharacterDataBase should initialized DB with empty map")
}

func TestCharacterDataBase_GetAllCharacters(t *testing.T) {
	characterDb := initDB()

	characters := characterDb.GetAll()

	assert.Len(t, characters, 1, "Should return list of Characters")
}

func TestCharacterDataBase_GetMovieById_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.GetById(1000)

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_GetMovieById(t *testing.T) {
	characterDb := initDB()

	character, err := characterDb.GetById(1)

	expected := Character{
		Id:      1,
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}

	assert.Nil(t, err, "No error")
	assert.Equal(t, expected, character, "Should return Character for ID 1")
}

func TestCharacterDataBase_AddCharacter(t *testing.T) {
	characterDb := initDB()

	newCharacterId := characterDb.Add(createOtherCharacter())
	assert.Greater(t, newCharacterId, uint64(0), "Should add new character to db and return its ID")
	assert.Len(t, characterDb.db, 2, "Should be 2 characters in DB")
}

func TestCharacterDataBase_UpdateCharacter_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.Update(1000, createOtherCharacter())

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_UpdateCharacter(t *testing.T) {
	characterDb := initDB()

	characterUpdated, err := characterDb.Update(1, createOtherCharacter())

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, characterUpdated, "Should update Character for ID 1 and return true")
	assert.Equal(t, characterDb.db[1].Id, uint64(1), "ID should stay the same")
	assert.Equal(t, characterDb.db[1].Name, OTHER_CHARACTER, "Name should be updated")
	assert.Equal(t, characterDb.db[1].MovieId, OTHER_MOVIE_ID, "MovieId should be updated")
}

func TestCharacterDataBase_DeleteCharacter_IdNotFound(t *testing.T) {
	characterDb := initDB()

	_, err := characterDb.Delete(1000)

	assert.EqualError(t, err, "Character entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestCharacterDataBase_DeleteCharacter(t *testing.T) {
	characterDb := initDB()

	characterDeleted, err := characterDb.Delete(1)

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, characterDeleted, "Should deleted Character for ID 1 and return true")
	assert.Len(t, characterDb.db, 0, "DB should be empty after deleting character")
}
