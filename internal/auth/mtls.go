package auth

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func LoadTLSConfig(certFile, keyFile, caFile string, isServer bool) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("loading keypair failed: %w", err)
	}

	caCert, err := os.ReadFile(caFile);
	if err != nil {
		return nil, fmt.Errorf("reading CA file failed: %w", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs: caPool
	}

	if isServer {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
	} else {
		cfg.RootCAs = caPool
	}

	return cfg, nil
}
