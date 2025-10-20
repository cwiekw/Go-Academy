package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	rand2 "math/rand/v2"
	"os"
	"time"
)

type CertManager interface {
	GenerateCertificateBasedOnCACert(name string) (*x509.Certificate, *rsa.PrivateKey, error)
	GenerateCertificateBasedOnCert(name string, inCert *x509.Certificate, inKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error)
}

type CertManagerImpl struct {
	cacert     *x509.Certificate
	privateKey *rsa.PrivateKey
}

func New() *CertManagerImpl {

	catls, err := tls.LoadX509KeyPair(os.Getenv("GMM_CERT_PATH"), os.Getenv("GMM_KEY_PATH"))

	if err != nil {
		panic(err)
	}

	ca, err := x509.ParseCertificate(catls.Certificate[0])

	if err != nil {
		panic(err)
	}

	return &CertManagerImpl{cacert: ca, privateKey: catls.PrivateKey.(*rsa.PrivateKey)}
}

func (cm CertManagerImpl) GenerateCertificateBasedOnCACert(name string) (*x509.Certificate, *rsa.PrivateKey, error) {
	return cm.generateCertificatesBasedOnCert(name, cm.cacert, cm.privateKey)
}

func (cm CertManagerImpl) GenerateCertificateBasedOnCert(name string, inCert *x509.Certificate, inKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	return cm.generateCertificatesBasedOnCert(name, inCert, inKey)
}

func (cm CertManagerImpl) generateCertificatesBasedOnCert(name string, inCert *x509.Certificate, inKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	capb := x509.MarshalPKCS1PublicKey(inCert.PublicKey.(*rsa.PublicKey))

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(rand2.Int64()),
		Subject: pkix.Name{
			Organization: []string{"TTPSC"},
			Country:      []string{"PL"},
			Locality:     []string{"Kielce"},
			CommonName:   name,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		SubjectKeyId: []byte{capb[17], capb[29], capb[47], capb[73], capb[97]},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	cert_b, err := x509.CreateCertificate(rand.Reader, cert, inCert, pub, inKey)

	if err != nil {
		return nil, nil, err
	}

	parsedCert, err := x509.ParseCertificate(cert_b)
	if err != nil {
		return nil, nil, err
	}

	return parsedCert, priv, nil
}
