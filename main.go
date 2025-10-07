package main

import (
	"MovieManager/character"
	"MovieManager/movie"
	"fmt"
)

func main() {

	movieDb := movie.NewMovieDataBase()
	characterDb := character.NewCharacterDataBase()

	m1 := movieDb.Add(movie.NewMovie(movie.Config{
		Name: "Matrix",
		Year: 1999,
	}))

	characterDb.Add(character.NewCharacter(character.Config{
		Name:    "Neo",
		MovieId: m1,
	}))

	characterDb.Add(character.NewCharacter(character.Config{
		Name:    "Agent Smith",
		MovieId: m1,
	}))

	m2 := movieDb.Add(movie.NewMovie(movie.Config{
		Name: "Casino Rolaye",
		Year: 2002,
	}))

	c1 := characterDb.Add(character.NewCharacter(character.Config{
		Name:    "Bames Jond",
		MovieId: m2,
	}))

	c2 := characterDb.Add(character.NewCharacter(character.Config{
		Name:    "N",
		MovieId: 2,
	}))

	m3 := movieDb.Add(movie.NewMovie(movie.Config{
		Name: "Mandalorian",
		Year: 2019,
	}))

	c3 := characterDb.Add(character.NewCharacter(character.Config{
		Name:    "Mando",
		MovieId: m3,
	}))

	_, _ = movieDb.Update(m2, movie.NewMovie(movie.Config{
		Name: "Casino Royale",
		Year: 2006,
	}))

	_, _ = characterDb.Update(c1, character.NewCharacter(character.Config{
		Name:    "James Bond",
		MovieId: m2,
	}))

	_, _ = characterDb.Update(c2, character.NewCharacter(character.Config{
		Name:    "M",
		MovieId: m2,
	}))

	_, _ = movieDb.Delete(m3)
	_, _ = characterDb.Delete(c3)

	allMovies := movieDb.GetAll()
	fmt.Println("Movies:")
	for _, movieEntry := range allMovies {
		fmt.Println(movieEntry)
	}

	fmt.Println()
	fmt.Println()

	allCharacters := characterDb.GetAll()
	fmt.Println("Characters:")
	for _, characterEntry := range allCharacters {
		movieForCharacter, _ := movieDb.GetById(characterEntry.MovieId)
		fmt.Printf("%s: %s\n", characterEntry.Name, movieForCharacter.Name)
	}
}
