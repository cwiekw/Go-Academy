package movie

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

type Movie struct {
	Id   uint64
	Name string
	Year uint16
	cert *x509.Certificate
	key  *rsa.PrivateKey
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

func WithCert(cert *x509.Certificate, key *rsa.PrivateKey) func(*Movie) {
	return func(m *Movie) {
		m.cert = cert
		m.key = key
	}
}

func (m Movie) GetCert() *x509.Certificate {
	return m.cert
}

func (m Movie) GetKey() *rsa.PrivateKey {
	return m.key
}

func (m Movie) String() string {
	return fmt.Sprintf("%s | %d", m.Name, m.Year)
}
