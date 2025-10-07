package movie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const OTHER_MOVIE = "Other Movie"
const OTHER_YEAR = uint16(1993)

func movieDbInitialized() MovieDataBase {
	return MovieDataBase{
		db: map[uint64]Movie{
			1: {
				Id:   1,
				Name: MOVIE,
				Year: YEAR,
			},
		},
	}
}

func movieDbNotInitialized() MovieDataBase {
	return MovieDataBase{}
}

func createOtherMovie() Movie {
	return Movie{
		Name: OTHER_MOVIE,
		Year: OTHER_YEAR,
	}
}

func TestNewMovieDataBase(t *testing.T) {
	result := NewMovieDataBase()

	expected := make(map[uint64]Movie)

	assert.Equal(t, expected, result.db, "NewMovieDataBase should initialize DB with empty map")
}

func TestMovieDataBaseCreateWithoutNew(t *testing.T) {
	result := MovieDataBase{}

	assert.Nil(t, result.db, "DB should be nil when creating directly")
}

func TestMovieDataBase_GetAllMovies_DBNotInitialized(t *testing.T) {
	movieDb := movieDbNotInitialized()

	_, err := movieDb.GetAllMovies()

	assert.EqualError(t, err, "Movie database not initialized!", "Should return error for not initialized DB")
}

func TestMovieDataBase_GetAllMovies(t *testing.T) {
	movieDb := movieDbInitialized()

	movies, err := movieDb.GetAllMovies()

	assert.Nil(t, err, "No error")
	assert.Len(t, movies, 1, "Should return list of Movies")
}

func TestMovieDataBase_GetMovieById_DBNotInitialized(t *testing.T) {
	movieDb := movieDbNotInitialized()

	_, err := movieDb.GetMovieById(1)

	assert.EqualError(t, err, "Movie database not initialized!", "Should return error for not initialized DB")
}

func TestMovieDataBase_GetMovieById_IdNotFound(t *testing.T) {
	movieDb := movieDbInitialized()

	_, err := movieDb.GetMovieById(1000)

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_GetMovieById(t *testing.T) {
	movieDb := movieDbInitialized()

	movie, err := movieDb.GetMovieById(1)

	expected := Movie{
		Id:   1,
		Name: MOVIE,
		Year: YEAR,
	}

	assert.Nil(t, err, "No error")
	assert.Equal(t, expected, movie, "Should return Movie for ID 1")
}

func TestMovieDataBase_AddMovie_DBNotInitialized(t *testing.T) {
	movieDb := movieDbNotInitialized()

	_, err := movieDb.AddMovie(createOtherMovie())

	assert.EqualError(t, err, "Movie database not initialized!", "Should return error for not initialized DB")
}

func TestMovieDataBase_AddMovie(t *testing.T) {
	movieDb := movieDbInitialized()

	newMovieId, err := movieDb.AddMovie(createOtherMovie())

	assert.Nil(t, err, "No error")
	assert.Greater(t, newMovieId, uint64(0), "Should add new movie to db and return its ID")
	assert.Len(t, movieDb.db, 2, "Should be 2 movies in DB")
}

func TestMovieDataBase_UpdateMovie_DBNotInitialized(t *testing.T) {
	movieDb := movieDbNotInitialized()

	_, err := movieDb.UpdateMovie(1, createOtherMovie())

	assert.EqualError(t, err, "Movie database not initialized!", "Should return error for not initialized DB")
}

func TestMovieDataBase_UpdateMovie_IdNotFound(t *testing.T) {
	movieDb := movieDbInitialized()

	_, err := movieDb.UpdateMovie(1000, createOtherMovie())

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_UpdateMovie(t *testing.T) {
	movieDb := movieDbInitialized()

	movieUpdated, err := movieDb.UpdateMovie(1, createOtherMovie())

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, movieUpdated, "Should update Movie for ID 1 and return true")
	assert.Equal(t, movieDb.db[1].Id, uint64(1), "ID should stay the same")
	assert.Equal(t, movieDb.db[1].Name, OTHER_MOVIE, "Name should be updated")
	assert.Equal(t, movieDb.db[1].Year, OTHER_YEAR, "Year should be updated")
}

func TestMovieDataBase_DeleteMovie_DBNotInitialized(t *testing.T) {
	movieDb := movieDbNotInitialized()

	_, err := movieDb.DeleteMovie(1)

	assert.EqualError(t, err, "Movie database not initialized!", "Should return error for not initialized DB")
}

func TestMovieDataBase_DeleteMovie_IdNotFound(t *testing.T) {
	movieDb := movieDbInitialized()

	_, err := movieDb.DeleteMovie(1000)

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_DeleteMovie(t *testing.T) {
	movieDb := movieDbInitialized()

	movieDeleted, err := movieDb.DeleteMovie(1)

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, movieDeleted, "Should delete Movie for ID 1 and return true")
	assert.Len(t, movieDb.db, 0, "DB should be empty after deleting movie")
}
