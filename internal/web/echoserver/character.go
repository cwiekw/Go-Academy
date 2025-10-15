package echoserver

import (
	"MovieManager/internal/web/api"
	"MovieManager/internal/web/mapper"
	"context"

	"go.uber.org/zap"
)

func (s Server) GetCharacters(_ context.Context, _ api.GetCharactersRequestObject) (api.GetCharactersResponseObject, error) {
	s.log.Info("GetCharacters", zap.String("action", "started"))
	characters := s.characterDb.GetAll()

	res := make([]api.Character, len(characters))

	for idx, c := range characters {
		res[idx] = mapper.MapCharacterEntityToDto(c)
	}

	s.log.Info("GetCharacters", zap.String("action", "finished"))
	return api.GetCharacters200JSONResponse(res), nil
}

func (s Server) GetCharactersCharacterId(_ context.Context, request api.GetCharactersCharacterIdRequestObject) (api.GetCharactersCharacterIdResponseObject, error) {
	s.log.Info("GetCharactersCharacterId", zap.String("action", "started"))
	c, err := s.characterDb.GetById(request.CharacterId)
	if err != nil {
		s.log.Info("GetCharactersCharacterId", zap.String("action", "failed"), zap.Error(err))
		return api.GetCharactersCharacterId404Response{}, nil
	}

	s.log.Info("GetCharactersCharacterId", zap.String("action", "finished"))
	return api.GetCharactersCharacterId200JSONResponse(mapper.MapCharacterEntityToDto(c)), nil
}

func (s Server) PostCharacters(_ context.Context, request api.PostCharactersRequestObject) (api.PostCharactersResponseObject, error) {
	s.log.Info("PostCharacters", zap.String("action", "started"))

	c := mapper.MapCharacterDtoToEntity(*request.Body)
	validated, err := s.validatorManager.Validate(c.MovieId, c.Name)

	if err != nil {
		s.log.Error("PostCharacters", zap.String("action", "validation failed"), zap.Error(err))
		return api.PostCharacters500Response{}, err
	}

	if !validated {
		s.log.Info("PostCharacters", zap.String("action", "validation does not allow character for movie. Breaking"))
		return api.PostCharacters412Response{}, nil
	}

	cm := s.characterDb.Add(c)

	s.log.Info("PostCharacters", zap.String("action", "finished"))
	return api.PostCharacters201JSONResponse(mapper.MapCharacterEntityToDto(cm)), nil
}

func (s Server) PutCharactersCharacterId(_ context.Context, request api.PutCharactersCharacterIdRequestObject) (api.PutCharactersCharacterIdResponseObject, error) {
	s.log.Info("PutCharactersCharacterId", zap.String("action", "started"))

	_, err := s.characterDb.Update(request.CharacterId, mapper.MapCharacterDtoToEntity(*request.Body))
	if err != nil {
		s.log.Info("PutCharactersCharacterId", zap.String("action", "failed"), zap.Error(err))
		return api.PutCharactersCharacterId404Response{}, nil
	}

	s.log.Info("PutCharactersCharacterId", zap.String("action", "finished"))
	return api.PutCharactersCharacterId204Response{}, nil
}

func (s Server) DeleteCharactersCharacterId(_ context.Context, request api.DeleteCharactersCharacterIdRequestObject) (api.DeleteCharactersCharacterIdResponseObject, error) {
	s.log.Info("DeleteCharactersCharacterId", zap.String("action", "started"))

	_, err := s.characterDb.Delete(request.CharacterId)
	if err != nil {
		s.log.Info("DeleteCharactersCharacterId", zap.String("action", "failed"), zap.Error(err))
		return api.DeleteCharactersCharacterId404Response{}, nil
	}

	s.log.Info("DeleteCharactersCharacterId", zap.String("action", "finished"))
	return api.DeleteCharactersCharacterId204Response{}, nil
}
