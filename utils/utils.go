package utils

import (
	"fmt"
	"math/rand/v2"
)

func GenerateId() uint64 {
	return rand.Uint64()
}

type DBNotInitializedError struct {
	dbName string
}

func (e *DBNotInitializedError) Error() string {
	return fmt.Sprintf("%s database not initialized!", e.dbName)
}

func NewDBNotInitializedError(dbName string) error {
	return &DBNotInitializedError{
		dbName: dbName,
	}
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
