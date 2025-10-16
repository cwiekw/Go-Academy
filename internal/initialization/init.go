package initialization

import (
	"MovieManager/internal/cert"
	"MovieManager/internal/database"
	"MovieManager/internal/entity/movie"
	"MovieManager/internal/validator"
)

func Init(mdb database.MovieDataBase, vm *validator.CharacterValidatorManager, cm cert.CertManager) {
	crt, key, err := cm.GenerateCertificateBasedOnCACert()
	if err != nil {
		return
	}

	starWars := mdb.Add(movie.New(
		movie.WithName("Star Wars"),
		movie.WithYear(1977),
		movie.WithCert(crt, key),
	))
	vm.AddValidator(starWars.Id, validator.NewStarWarsCharacterValidator())
}
