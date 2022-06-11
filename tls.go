package web

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
)

// TLS13Config returns *tls.Config that only supports TLS 1.3 and does not
// enforce client side authentication.
func TLS13Config() *tls.Config {
	return &tls.Config{
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS13SecureCipherList(),
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
	}
}

// MutualTLS13Config returns *tls.Config that supports mTLS 1.3 with
// verification via the specified certificate authority.
func MutualTLS13Config(caCertPool *x509.CertPool) *tls.Config {
	return &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                caCertPool,
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS13SecureCipherList(),
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
	}
}

// TLS13SecureCipherList returns a slice of TLS 1.3 cipher suites.
func TLS13SecureCipherList() []uint16 {
	return []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
	}
}

// TLS12Config returns *tls.Config with TLS 1.2 settings.
func TLS12Config() *tls.Config {
	return &tls.Config{
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS12SecureCipherList(),
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}
}

// MutualTLS12Config returns *tls.Config that supports mTLS 1.2 with
// verification via the specified certificate authority.
func MutualTLS12Config(caCertPool *x509.CertPool) *tls.Config {
	return &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                caCertPool,
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS12SecureCipherList(),
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}
}

// TLS12SecureCipherList returns a slice of TLS 1.2 cipher suites.
func TLS12SecureCipherList() []uint16 {
	return []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	}
}

// LoadCACerts parses cacertFile for cert files and adds them to a CertPool
// caCertPool, returning caCertPool with nil error if successful.
// Else, return nil CertPool and the error encountered.
func LoadCACerts(cacertFile string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(cacertFile)
	if err != nil {
		return nil, fmt.Errorf("error loading cert file: %w", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return caCertPool, nil
}

type TLS struct {
	CACertFile  string `json:"caCertFile"`
	CertFile    string `json:"certFile"`
	KeyFile     string `json:"keyFile"`
	EnableTLS12 bool   `json:"tls12,omitempty"`
}

func (tc *TLS) Listener(addr *Address) (net.Listener, error) {
	if tc == nil {
		return addr.Listener()
	}

	cert, err := tls.LoadX509KeyPair(tc.CertFile, tc.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("cert load error: %w", err)
	}

	var cfg *tls.Config

	if tc.CACertFile != "" {
		// CA cert is present means we should do MTLS
		caCerts, err := LoadCACerts(tc.CACertFile)
		if err != nil {
			return nil, err
		}

		if tc.EnableTLS12 {
			cfg = MutualTLS12Config(caCerts)
		} else {
			cfg = MutualTLS13Config(caCerts)
		}
	} else {
		if tc.EnableTLS12 {
			cfg = TLS12Config()
		} else {
			cfg = TLS13Config()
		}
	}

	// Set the cert
	cfg.Certificates = []tls.Certificate{cert}

	l, err := addr.Listener()
	if err != nil {
		return nil, err
	}

	return tls.NewListener(l, cfg), nil
}
