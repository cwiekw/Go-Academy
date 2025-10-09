package character

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const CHARACTER = "Character"
const MOVIE_ID = uint64(12345)

func TestNew_NoParams(t *testing.T) {
	m := New()

	expected := Character{
		Id:      0,
		Name:    "",
		MovieId: 0,
	}

	assert.Equal(t, expected, m, "Should return Character with all zero values")
}

func TestNew_WithName(t *testing.T) {
	m := New(
		WithName(CHARACTER),
	)

	expected := Character{
		Id:      0,
		Name:    CHARACTER,
		MovieId: 0,
	}

	assert.Equal(t, expected, m, "Should return Character with Name")
}

func TestNew_WithMovieId(t *testing.T) {
	m := New(
		WithMovieId(MOVIE_ID),
	)

	expected := Character{
		Id:      0,
		Name:    "",
		MovieId: MOVIE_ID,
	}

	assert.Equal(t, expected, m, "Should return Character with MovieId")
}

func TestNew_WithNameAndMovieId(t *testing.T) {
	m := New(
		WithName(CHARACTER),
		WithMovieId(MOVIE_ID),
	)

	expected := Character{
		Id:      0,
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}

	assert.Equal(t, expected, m, "Should return Character with Name and MovieId")
}

func TestCharacter_String(t *testing.T) {
	result := Character{
		Name:    CHARACTER,
		MovieId: MOVIE_ID,
	}.String()

	expected := fmt.Sprintf("%s: %d", CHARACTER, MOVIE_ID)

	assert.Equal(t, expected, result, "Character.String() should return \"Character: 12345\"")
}
