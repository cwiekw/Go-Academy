package main

import (
	"MovieManager/character"
	"MovieManager/movie"
	"fmt"
)

func main() {

	movieDb := movie.NewMovieDataBase()
	characterDb := character.NewCharacterDataBase()

	m1, _ := movieDb.AddMovie(movie.NewMovie(movie.Config{
		Name: "Matrix",
		Year: 1999,
	}))

	_, _ = characterDb.AddCharacter(character.NewCharacter(character.Config{
		Name:    "Neo",
		MovieId: m1,
	}))

	_, _ = characterDb.AddCharacter(character.NewCharacter(character.Config{
		Name:    "Agent Smith",
		MovieId: m1,
	}))

	m2, _ := movieDb.AddMovie(movie.NewMovie(movie.Config{
		Name: "Casino Rolaye",
		Year: 2002,
	}))

	c1, _ := characterDb.AddCharacter(character.NewCharacter(character.Config{
		Name:    "Bames Jond",
		MovieId: m2,
	}))

	c2, _ := characterDb.AddCharacter(character.NewCharacter(character.Config{
		Name:    "N",
		MovieId: 2,
	}))

	m3, _ := movieDb.AddMovie(movie.NewMovie(movie.Config{
		Name: "Mandalorian",
		Year: 2019,
	}))

	c3, _ := characterDb.AddCharacter(character.NewCharacter(character.Config{
		Name:    "Mando",
		MovieId: m3,
	}))

	_, _ = movieDb.UpdateMovie(m2, movie.NewMovie(movie.Config{
		Name: "Casino Royale",
		Year: 2006,
	}))

	_, _ = characterDb.UpdateCharacter(c1, character.NewCharacter(character.Config{
		Name:    "James Bond",
		MovieId: m2,
	}))

	_, _ = characterDb.UpdateCharacter(c2, character.NewCharacter(character.Config{
		Name:    "M",
		MovieId: m2,
	}))

	_, _ = movieDb.DeleteMovie(m3)
	_, _ = characterDb.DeleteCharacter(c3)

	allMovies, _ := movieDb.GetAllMovies()
	fmt.Println("Movies:")
	for _, movieEntry := range allMovies {
		fmt.Println(movieEntry)
	}

	fmt.Println()
	fmt.Println()

	allCharacters, _ := characterDb.GetAllCharacters()
	fmt.Println("Characters:")
	for _, characterEntry := range allCharacters {
		movieForCharacter, _ := movieDb.GetMovieById(characterEntry.MovieId)
		fmt.Printf("%s: %s\n", characterEntry.Name, movieForCharacter.Name)
	}
}
