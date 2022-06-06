package web

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func TLS13Config() *tls.Config {
	return &tls.Config{
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS13SecureCipherList(),
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
	}
}

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

func TLS13SecureCipherList() []uint16 {
	return []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
	}
}

func TLS12Config() *tls.Config {
	return &tls.Config{
		MaxVersion:               tls.VersionTLS13,
		CipherSuites:             TLS12SecureCipherList(),
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}
}

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

func LoadCACerts(cacertFile string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(cacertFile)
	if err != nil {
		return nil, fmt.Errorf("error loading cert file: %w", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return caCertPool, nil
}
