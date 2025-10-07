package character

import "fmt"

type Character struct {
	Id      uint64
	Name    string
	MovieId uint64
}

func New(options ...func(*Character)) Character {
	c := Character{}

	for _, o := range options {
		o(&c)
	}

	return c
}

func WithName(name string) func(*Character) {
	return func(c *Character) {
		c.Name = name
	}
}

func WithMovieId(movieId uint64) func(*Character) {
	return func(c *Character) {
		c.MovieId = movieId
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%s: %d", c.Name, c.MovieId)
}
