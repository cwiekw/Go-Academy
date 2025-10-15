package movie

import (
	"MovieManager/internal/entity/movie"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MOVIE = "Movie"
const YEAR = uint16(2017)
const OTHER_MOVIE = "Other Movie"
const OTHER_YEAR = uint16(1993)

func initDB() InMemoryMovieDataBase {
	return InMemoryMovieDataBase{
		db: map[uint64]movie.Movie{
			1: {
				Id:   1,
				Name: MOVIE,
				Year: YEAR,
			},
		},
	}
}

func createOtherMovie() movie.Movie {
	return movie.Movie{
		Name: OTHER_MOVIE,
		Year: OTHER_YEAR,
	}
}

func TestNewMovieDataBase(t *testing.T) {
	result := New()

	expected := make(map[uint64]movie.Movie)

	assert.Equal(t, expected, result.db, "New should initialize DB with empty map")
}

func TestMovieDataBaseCreateWithoutNew(t *testing.T) {
	result := InMemoryMovieDataBase{}

	assert.Nil(t, result.db, "DB should be nil when creating directly")
}

func TestMovieDataBase_GetAllMovies(t *testing.T) {
	movieDb := initDB()

	movies := movieDb.GetAll()

	assert.Len(t, movies, 1, "Should return list of Movies")
}

func TestMovieDataBase_GetMovieById_IdNotFound(t *testing.T) {
	movieDb := initDB()

	_, err := movieDb.GetById(1000)

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_GetMovieById(t *testing.T) {
	movieDb := initDB()

	m, err := movieDb.GetById(1)

	expected := movie.Movie{
		Id:   1,
		Name: MOVIE,
		Year: YEAR,
	}

	assert.Nil(t, err, "No error")
	assert.Equal(t, expected, m, "Should return Movie for ID 1")
}

func TestMovieDataBase_AddMovie(t *testing.T) {
	movieDb := initDB()

	newMovie := movieDb.Add(createOtherMovie())

	assert.Greater(t, newMovie.Id, uint64(0), "Should add new movie to db and return it. Generated ID should be greater than 0")
	assert.Equal(t, OTHER_MOVIE, newMovie.Name, "Should add new movie to db and return it. Name should be equal to provided one")
	assert.Equal(t, OTHER_YEAR, newMovie.Year, "Should add new movie to db and return it. Year should be equal to provided one")
	assert.Len(t, movieDb.db, 2, "Should be 2 movies in DB")
}

func TestMovieDataBase_UpdateMovie_IdNotFound(t *testing.T) {
	movieDb := initDB()

	_, err := movieDb.Update(1000, createOtherMovie())

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_UpdateMovie(t *testing.T) {
	movieDb := initDB()

	movieUpdated, err := movieDb.Update(1, createOtherMovie())

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, movieUpdated, "Should update Movie for ID 1 and return true")
	assert.Equal(t, movieDb.db[1].Id, uint64(1), "ID should stay the same")
	assert.Equal(t, movieDb.db[1].Name, OTHER_MOVIE, "Name should be updated")
	assert.Equal(t, movieDb.db[1].Year, OTHER_YEAR, "Year should be updated")
}

func TestMovieDataBase_DeleteMovie_IdNotFound(t *testing.T) {
	movieDb := initDB()

	_, err := movieDb.Delete(1000)

	assert.EqualError(t, err, "Movie entity for id 1000 does not exist!", "Should return error for not found ID")
}

func TestMovieDataBase_DeleteMovie(t *testing.T) {
	movieDb := initDB()

	movieDeleted, err := movieDb.Delete(1)

	assert.Nil(t, err, "No error")
	assert.Equal(t, true, movieDeleted, "Should delete Movie for ID 1 and return true")
	assert.Len(t, movieDb.db, 0, "DB should be empty after deleting movie")
}
