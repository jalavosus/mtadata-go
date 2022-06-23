package serverauth

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jalavosus/mtadata/internal/config"
)

var Module = fx.Options(fx.Provide(NewServerAuth))

type ServerAuth struct {
	privateKey *ecdsa.PrivateKey
	x509Cert   *x509.Certificate
	caCert     *x509.Certificate
	certPool   *x509.CertPool
	useTls     bool
}

func NewServerAuth(tlsCfg *config.TlsConfig) (s *ServerAuth, err error) {
	s = &ServerAuth{
		certPool: x509.NewCertPool(),
		useTls:   tlsCfg.UseTls,
	}

	if !s.useTls {
		return
	}

	s.privateKey, err = parseKey(tlsCfg.KeyPath)
	if err != nil {
		s = nil
		return
	}

	s.x509Cert, err = parseCertificate(tlsCfg.CertPath)
	if err != nil {
		s = nil
		return
	}

	s.certPool.AddCert(s.x509Cert)

	if tlsCfg.CaPath != "" {
		s.caCert, err = parseCertificate(tlsCfg.CaPath)
		if err != nil {
			s = nil
			return
		}

		s.certPool.AddCert(s.caCert)
	}

	return
}

func (s *ServerAuth) TlsCert() *tls.Certificate {
	return x509CertToTls(s.privateKey, s.x509Cert)
}

func (s *ServerAuth) CACert() *tls.Certificate {
	return x509CertToTls(s.privateKey, s.caCert)
}

func (s *ServerAuth) CertPool() *x509.CertPool {
	return s.certPool
}

func (s *ServerAuth) Key() crypto.PrivateKey {
	return s.privateKey
}

func (s *ServerAuth) TransportCredentials(client bool) credentials.TransportCredentials {
	if s.useTls {
		if client {
			return credentials.NewClientTLSFromCert(s.certPool, "")
		}

		return credentials.NewServerTLSFromCert(s.TlsCert())
	}

	return insecure.NewCredentials()
}

func parseCertificate(certFile string) (*x509.Certificate, error) {
	certBytes, err := readFileBytes(certFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certBytes)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.WithMessagef(err, "error parsing certificate data from %[1]s", certFile)
	}

	return cert, nil
}

func parseKey(keyFile string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := readFileBytes(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.WithMessagef(err, "error parsing key data from %[1]s", keyFile)
	}

	return key, nil
}

func readFileBytes(fp string) ([]byte, error) {
	if !filepath.IsAbs(fp) {
		var err error

		fp, err = filepath.Abs(fp)
		if err != nil {
			return nil, err
		}
	}

	return os.ReadFile(fp)
}

func x509CertToTls(privKey crypto.PrivateKey, cert *x509.Certificate) *tls.Certificate {
	rawCerts := [][]byte{cert.Raw}

	return &tls.Certificate{
		Certificate: rawCerts,
		PrivateKey:  privKey,
		Leaf:        cert,
	}
}