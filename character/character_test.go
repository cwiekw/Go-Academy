package character

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const CHARACTER = "Character"
const MOVIE_ID = uint64(12345)

func TestNewCharacter(t *testing.T) {
	result := NewCharacter(Config{
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	})

	expected := Character{
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}

	assert.Equal(t, expected, result, "NewCharacter should create new Movie entity")
}

func TestCharacter_String(t *testing.T) {
	result := NewCharacter(Config{
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}).String()

	expected := fmt.Sprintf("%s: %d", CHARACTER, MOVIE_ID)

	assert.Equal(t, expected, result, "Character.String() should return \"Character: 12345\"")
}
