package movie

import (
	"fmt"
)

type Config struct {
	Name string
	Year uint16
}

type Movie struct {
	Id   uint64
	Name string
	Year uint16
}

func NewMovie(cfg Config) Movie {
	return Movie{
		Name: cfg.Name,
		Year: cfg.Year,
	}
}

func (m Movie) String() string {
	return fmt.Sprintf("%s | %d", m.Name, m.Year)
}
