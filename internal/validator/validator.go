package validator

import (
	"encoding/json"
	"fmt"
	"strings"

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
	url := strings.ReplaceAll(fmt.Sprintf("https://swapi.dev/api/people/?search=%s", name), " ", "%20")

	response, err := c.httpClient.R().Get(url)

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
