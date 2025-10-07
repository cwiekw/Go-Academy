package character

import "fmt"

type Config struct {
	Name    string
	MovieId uint64
}

type Character struct {
	Id      uint64
	Name    string
	MovieId uint64
}

func NewCharacter(cfg Config) Character {
	return Character{
		Name:    cfg.Name,
		MovieId: cfg.MovieId,
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%s: %d", c.Name, c.MovieId)
}
