package utils

import (
	"fmt"
	"math/rand/v2"
)

func GenerateId() uint64 {
	return rand.Uint64()
}

type EntityDoesNotExistError struct {
	entityType string
	id         uint64
}

func (e *EntityDoesNotExistError) Error() string {
	return fmt.Sprintf("%s entity for id %d does not exist!", e.entityType, e.id)
}

func NewEntityDoesNotExistError(entityType string, id uint64) error {
	return &EntityDoesNotExistError{
		entityType: entityType,
		id:         id,
	}
}
