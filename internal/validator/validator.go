package validator

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type CharacterValidator interface {
	Validate(name string) (bool, error)
}

type StarWarsCharacterValidator struct {
	httpClient *resty.Client
}

type starWarsResult struct {
	Count int `json:"count"`
}

func NewStarWarsCharacterValidator() *StarWarsCharacterValidator {
	return &StarWarsCharacterValidator{
		httpClient: resty.New(),
	}
}

func (c StarWarsCharacterValidator) Validate(name string) (bool, error) {
	response, err := c.httpClient.R().
		SetQueryParam("search", name).
		Get("https://swapi.dev/api/people")

	if err != nil {
		return false, err
	}

	result := starWarsResult{}
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return false, err
	}

	if result.Count == 0 {
		return false, nil
	}

	return true, nil
}
