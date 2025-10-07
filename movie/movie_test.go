package movie

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

const MOVIE = "Movie"
const YEAR = uint16(2017)

func TestNewMovie(t *testing.T) {
	result := NewMovie(Config{
		Name: MOVIE,
		Year: YEAR,
	})

	expected := Movie{
		Name: MOVIE,
		Year: YEAR,
	}

	assert.Equal(t, expected, result, "NewMovie should create new Movie entity")
}

func TestMovie_String(t *testing.T) {
	result := NewMovie(Config{
		Name: MOVIE,
		Year: YEAR,
	}).String()

	expected := fmt.Sprintf("%s | %d", MOVIE, YEAR)

	assert.Equal(t, expected, result, "Movie.String() should return \"Movie | 2017\"")
}
