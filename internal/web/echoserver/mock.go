package echoserver

import (
	"MovieManager/internal/database"
	echaracter "MovieManager/internal/entity/character"
	emovie "MovieManager/internal/entity/movie"
	"MovieManager/internal/validator"
	"crypto/rsa"
	"crypto/x509"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

var ID = uint64(1)
var WRONG_ID = uint64(999)
var MOVIE_ID = uint64(1001)
var CRT = &x509.Certificate{}
var KEY = &rsa.PrivateKey{}

func newServer() *Server {
	return New(zap.NewNop(), newMockMovieDataBase(), newMockCharacterDataBase(), newCharacterValidatorManager(), newMockCertManager())
}

type mockMovieDataBase struct {
	mock.Mock
}

func newMockMovieDataBase() *mockMovieDataBase {
	m := &mockMovieDataBase{}

	movie := emovie.New(
		emovie.WithName("Movie1"),
		emovie.WithYear(2001),
		emovie.WithCert(CRT, KEY),
	)

	simpleMovie := emovie.New(
		emovie.WithName("Movie1"),
		emovie.WithYear(2001),
	)

	movieWithId := emovie.New(
		emovie.WithName("Movie1"),
		emovie.WithYear(2001),
		emovie.WithCert(CRT, KEY),
	)

	movieWithId.Id = ID

	m.On("GetAll").Return([]emovie.Movie{
		movieWithId,
		{
			Id:   WRONG_ID,
			Name: "Movie2",
			Year: 2002,
		},
	})

	m.On("Add", movie).Return(movieWithId)

	m.On("GetById", ID).Return(movieWithId, nil)
	m.On("GetById", MOVIE_ID).Return(movieWithId, nil)
	m.On("GetById", WRONG_ID).Return(emovie.Movie{}, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

	m.On("Update", ID, simpleMovie).Return(true, nil)
	m.On("Update", WRONG_ID, simpleMovie).Return(false, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

	m.On("Delete", ID).Return(true, nil)
	m.On("Delete", WRONG_ID).Return(false, database.NewEntityDoesNotExistError("Movie", WRONG_ID))

	return m
}

func (m *mockMovieDataBase) GetAll() []emovie.Movie {
	args := m.Called()
	return args.Get(0).([]emovie.Movie)
}

func (m *mockMovieDataBase) GetById(id uint64) (emovie.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(emovie.Movie), args.Error(1)
}

func (m *mockMovieDataBase) Add(em emovie.Movie) emovie.Movie {
	args := m.Called(em)
	return args.Get(0).(emovie.Movie)
}

func (m *mockMovieDataBase) Update(id uint64, u emovie.Movie) (bool, error) {
	args := m.Called(id, u)
	return args.Bool(0), args.Error(1)
}

func (m *mockMovieDataBase) Delete(id uint64) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

type mockCharacterDataBase struct {
	mock.Mock
}

func newMockCharacterDataBase() *mockCharacterDataBase {
	m := &mockCharacterDataBase{}

	character := echaracter.New(
		echaracter.WithName("Character1"),
		echaracter.WithMovieId(MOVIE_ID),
		echaracter.WithCert(CRT, KEY),
	)

	simpleCharacter := echaracter.New(
		echaracter.WithName("Character1"),
		echaracter.WithMovieId(MOVIE_ID),
	)

	characterWithId := echaracter.New(
		echaracter.WithName("Character1"),
		echaracter.WithMovieId(MOVIE_ID),
		echaracter.WithCert(CRT, KEY),
	)

	characterWithId.Id = ID

	m.On("GetAll").Return([]echaracter.Character{
		characterWithId,
		{
			Id:      WRONG_ID,
			Name:    "Character2",
			MovieId: uint64(1002),
		},
	})

	m.On("Add", character).Return(characterWithId)

	m.On("GetById", ID).Return(characterWithId, nil)
	m.On("GetById", WRONG_ID).Return(echaracter.Character{}, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	m.On("Update", ID, simpleCharacter).Return(true, nil)
	m.On("Update", WRONG_ID, simpleCharacter).Return(false, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	m.On("Delete", ID).Return(true, nil)
	m.On("Delete", WRONG_ID).Return(false, database.NewEntityDoesNotExistError("Character", WRONG_ID))

	return m
}

func (m *mockCharacterDataBase) GetAll() []echaracter.Character {
	args := m.Called()
	return args.Get(0).([]echaracter.Character)
}

func (m *mockCharacterDataBase) GetById(id uint64) (echaracter.Character, error) {
	args := m.Called(id)
	return args.Get(0).(echaracter.Character), args.Error(1)
}

func (m *mockCharacterDataBase) Add(em echaracter.Character) echaracter.Character {
	args := m.Called(em)
	return args.Get(0).(echaracter.Character)
}

func (m *mockCharacterDataBase) Update(id uint64, u echaracter.Character) (bool, error) {
	args := m.Called(id, u)
	return args.Bool(0), args.Error(1)
}

func (m *mockCharacterDataBase) Delete(id uint64) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

type mockCharacterValidator struct {
	mock.Mock
}

func newMockCharacterValidator() *mockCharacterValidator {
	v := &mockCharacterValidator{}

	v.On("Validate", "Character1").Return(true, nil)
	v.On("Validate", "Character2").Return(false, nil)
	v.On("Validate", "Character3").Return(false, errors.New("E"))

	return v
}

func (m *mockCharacterValidator) Validate(name string) (bool, error) {
	args := m.Called(name)
	return args.Bool(0), args.Error(1)
}

func newCharacterValidatorManager() validator.CharacterValidatorManager {
	vm := validator.NewCharacterValidatorManager()
	vm.AddValidator(MOVIE_ID, newMockCharacterValidator())

	return *vm
}

type mockCertManager struct {
	mock.Mock
}

func (m *mockCertManager) GenerateCertificateBasedOnCACert(name string) (*x509.Certificate, *rsa.PrivateKey, error) {
	args := m.Called(name)
	return args.Get(0).(*x509.Certificate), args.Get(1).(*rsa.PrivateKey), args.Error(2)
}

func (m *mockCertManager) GenerateCertificateBasedOnCert(name string, inCert *x509.Certificate, inKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	args := m.Called(name, inCert, inKey)
	return args.Get(0).(*x509.Certificate), args.Get(1).(*rsa.PrivateKey), args.Error(2)
}

func newMockCertManager() *mockCertManager {
	m := &mockCertManager{}

	m.On("GenerateCertificateBasedOnCACert", mock.Anything).Return(CRT, KEY, nil)
	m.On("GenerateCertificateBasedOnCert", mock.Anything, mock.Anything, mock.Anything).Return(CRT, KEY, nil)

	return m
}
