package initialization

import (
	"MovieManager/internal/database"
	"MovieManager/internal/entity/movie"
	"MovieManager/internal/validator"
)

func Init(mdb database.MovieDataBase, vm *validator.CharacterValidatorManager) {
	starWars := mdb.Add(movie.Movie{Name: "Star Wars", Year: 1977})
	vm.AddValidator(starWars.Id, validator.NewStarWarsCharacterValidator())
}
