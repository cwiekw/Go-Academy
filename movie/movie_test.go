package movie

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

const MOVIE = "Movie"
const YEAR = uint16(2017)

func TestNew_NoParams(t *testing.T) {
	m := New()

	expected := Movie{
		Id:   0,
		Name: "",
		Year: 0,
	}

	assert.Equal(t, expected, m, "Should return Movie with all zero values")
}

func TestNew_WithName(t *testing.T) {
	m := New(
		WithName(MOVIE),
	)

	expected := Movie{
		Id:   0,
		Name: MOVIE,
		Year: 0,
	}

	assert.Equal(t, expected, m, "Should return Movie with Name")
}

func TestNew_WithYear(t *testing.T) {
	m := New(
		WithYear(YEAR),
	)

	expected := Movie{
		Id:   0,
		Name: "",
		Year: YEAR,
	}

	assert.Equal(t, expected, m, "Should return Movie with Year")
}

func TestNew_WithNameAndYear(t *testing.T) {
	m := New(
		WithName(MOVIE),
		WithYear(YEAR),
	)

	expected := Movie{
		Id:   0,
		Name: MOVIE,
		Year: YEAR,
	}

	assert.Equal(t, expected, m, "Should return Movie with Name and Year")
}

func TestMovie_String(t *testing.T) {
	result := Movie{
		Name: MOVIE,
		Year: YEAR,
	}.String()

	expected := fmt.Sprintf("%s | %d", MOVIE, YEAR)

	assert.Equal(t, expected, result, "Movie.String() should return \"Movie | 2017\"")
}
