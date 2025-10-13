package echoserver

import (
	"MovieManager/internal/web/api"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_GetCharacters(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetCharacters(c, api.GetCharactersRequestObject{})

	expected := api.GetCharacters200JSONResponse([]api.Character{
		{
			Id:      &ID,
			Name:    "Character1",
			MovieId: MOVIE_ID,
		},
		{
			Id:      &WRONG_ID,
			Name:    "Character2",
			MovieId: uint64(1002),
		},
	})

	assert.Equal(t, expected, res, "Should return all characters")
}

func TestServer_GetCharactersCharacterId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetCharactersCharacterId(c, api.GetCharactersCharacterIdRequestObject{CharacterId: ID})

	expected := api.GetCharactersCharacterId200JSONResponse{
		Id:      &ID,
		Name:    "Character1",
		MovieId: MOVIE_ID,
	}

	assert.Equal(t, expected, res, "Should return character by ID")
}

func TestServer_GetCharactersCharacterId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.GetCharactersCharacterId(c, api.GetCharactersCharacterIdRequestObject{CharacterId: WRONG_ID})

	expected := api.GetCharactersCharacterId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}

func TestServer_PostCharacters(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PostCharacters(c, api.PostCharactersRequestObject{Body: &api.Character{
		Name:    "Character1",
		MovieId: MOVIE_ID,
	}})

	expected := api.PostCharacters201JSONResponse(api.Character{
		Id:      &ID,
		Name:    "Character1",
		MovieId: MOVIE_ID,
	})

	assert.Equal(t, expected, res, "Should add new character and return it with 201 entity")
}

func TestServer_PostCharacters_ValidationFailed(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PostCharacters(c, api.PostCharactersRequestObject{Body: &api.Character{
		Name:    "Character2",
		MovieId: MOVIE_ID,
	}})

	expected := api.PostCharacters412Response{}

	assert.Equal(t, expected, res, "Should fail validation and return 412 entity")
}

func TestServer_PostCharacters_ValidationError(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, err := s.PostCharacters(c, api.PostCharactersRequestObject{Body: &api.Character{
		Name:    "Character3",
		MovieId: MOVIE_ID,
	}})

	expected := api.PostCharacters500Response{}

	assert.Equal(t, expected, res, "Validation should throw error and return 500 entity")
	assert.Error(t, err, "Error should be there")
}

func TestServer_PutCharactersCharacterId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PutCharactersCharacterId(c, api.PutCharactersCharacterIdRequestObject{
		CharacterId: ID,
		Body: &api.Character{
			Name:    "Character1",
			MovieId: MOVIE_ID,
		}})

	expected := api.PutCharactersCharacterId204Response{}

	assert.Equal(t, expected, res, "Should update character and return 204 entity")
}

func TestServer_PutCharactersCharacterId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.PutCharactersCharacterId(c, api.PutCharactersCharacterIdRequestObject{
		CharacterId: WRONG_ID,
		Body: &api.Character{
			Name:    "Character1",
			MovieId: MOVIE_ID,
		}})

	expected := api.PutCharactersCharacterId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}

func TestServer_DeleteCharactersCharacterId(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.DeleteCharactersCharacterId(c, api.DeleteCharactersCharacterIdRequestObject{CharacterId: ID})

	expected := api.DeleteCharactersCharacterId204Response{}

	assert.Equal(t, expected, res, "Should delete character and return 204 entity")
}

func TestServer_DeleteCharactersCharacterId_NotFound(t *testing.T) {
	s := newServer()
	c := context.Background()

	res, _ := s.DeleteCharactersCharacterId(c, api.DeleteCharactersCharacterIdRequestObject{CharacterId: WRONG_ID})

	expected := api.DeleteCharactersCharacterId404Response{}

	assert.Equal(t, expected, res, "Should return 404 entity for not found ID")
}
