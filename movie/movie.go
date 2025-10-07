package movie

import (
	"fmt"
)

type Movie struct {
	Id   uint64
	Name string
	Year uint16
}

func New(options ...func(*Movie)) Movie {
	m := Movie{}

	for _, o := range options {
		o(&m)
	}

	return m
}

func WithName(name string) func(*Movie) {
	return func(m *Movie) {
		m.Name = name
	}
}

func WithYear(year uint16) func(*Movie) {
	return func(m *Movie) {
		m.Year = year
	}
}

func (m Movie) String() string {
	return fmt.Sprintf("%s | %d", m.Name, m.Year)
}
