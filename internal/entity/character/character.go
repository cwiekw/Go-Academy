package character

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

type Character struct {
	Id      uint64
	Name    string
	MovieId uint64
	cert    *x509.Certificate
	key     *rsa.PrivateKey
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

func WithCert(cert *x509.Certificate, key *rsa.PrivateKey) func(*Character) {
	return func(c *Character) {
		c.cert = cert
		c.key = key
	}
}

func (c Character) GetCert() *x509.Certificate {
	return c.cert
}

func (c Character) String() string {
	return fmt.Sprintf("%s: %d", c.Name, c.MovieId)
}
