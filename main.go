package main

import (
	"MovieManager/character"
	"MovieManager/movie"
	"fmt"
)

func main() {

	movieDb := movie.NewMovieDataBase()
	characterDb := character.NewCharacterDataBase()

	m1 := movieDb.Add(movie.New(
		movie.WithName("Matrix"),
		movie.WithYear(1999),
	))

	characterDb.Add(character.New(
		character.WithName("Neo"),
		character.WithMovieId(m1),
	))

	characterDb.Add(character.New(
		character.WithName("Agent Smith"),
		character.WithMovieId(m1),
	))

	m2 := movieDb.Add(movie.New(
		movie.WithName("Casino Rolaye"),
		movie.WithYear(2002),
	))

	c1 := characterDb.Add(character.New(
		character.WithName("Bames Jond"),
		character.WithMovieId(m2),
	))

	c2 := characterDb.Add(character.New(
		character.WithName("N"),
		character.WithMovieId(m2),
	))

	m3 := movieDb.Add(movie.New(
		movie.WithName("Mandalorian"),
		movie.WithYear(2019),
	))

	c3 := characterDb.Add(character.New(
		character.WithName("Mando"),
		character.WithMovieId(m3),
	))

	_, _ = movieDb.Update(m2, movie.New(
		movie.WithName("Casino Royale"),
		movie.WithYear(2006),
	))

	_, _ = characterDb.Update(c1, character.New(
		character.WithName("James Bond"),
		character.WithMovieId(m2),
	))

	_, _ = characterDb.Update(c2, character.New(
		character.WithName("M"),
		character.WithMovieId(m2),
	))

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
