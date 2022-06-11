package server_test

import (
	"testing"

	"gojini.dev/web"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	addr := web.NewAddress("127.0.0.1:7777")
	assert.NotNil(addr)
	assert.Equal(addr.String(), "127.0.0.1:7777")
	assert.Equal(addr.Format(), "tcp://127.0.0.1:7777")

	l, e := addr.Listener()
	assert.Nil(e)
	assert.NotNil(l)
	assert.Nil(l.Close())

	addr = web.NewAddress("gabbar.bad.address.com:8080")
	l, e = addr.Listener()
	assert.NotNil(e)
	assert.Nil(l)
}

func TestTLSListener(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	addr := web.NewAddress("127.0.0.1:6666")
	var tls *web.TLS
	l, e := tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())

	// Server TLS1.3
	tls = &web.TLS{
		CertFile: "../test_certs/server.crt",
		KeyFile:  "../test_certs/server.key",
	}
	l, e = tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())

	// Server TLS1.2
	tls = &web.TLS{
		CertFile:    "../test_certs/server.crt",
		KeyFile:     "../test_certs/server.key",
		EnableTLS12: true,
	}
	l, e = tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())
}

func TestMTLSListener(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	addr := web.NewAddress("127.0.0.1:5555")
	var tls *web.TLS
	l, e := tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())

	// Server MTLS1.3
	tls = &web.TLS{
		CACertFile: "../test_certs/ca.crt",
		CertFile:   "../test_certs/server.crt",
		KeyFile:    "../test_certs/server.key",
	}
	l, e = tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())

	// Server MTLS 1.2
	tls = &web.TLS{
		CACertFile:  "../test_certs/ca.crt",
		CertFile:    "../test_certs/server.crt",
		KeyFile:     "../test_certs/server.key",
		EnableTLS12: true,
	}
	l, e = tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())

	tls = &web.TLS{
		CertFile:    "../test_certs/server.crt",
		KeyFile:     "../test_certs/server.key",
		EnableTLS12: true,
	}
	l, e = tls.Listener(addr)
	assert.NotNil(l)
	assert.Nil(e)
	assert.Nil(l.Close())
}

func TestBadCerts(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	addr := web.NewAddress("127.0.0.1:5555")

	// Bad CA Cert
	tls := &web.TLS{
		CACertFile: "blah.crt",
		CertFile:   "../test_certs/server.crt",
		KeyFile:    "../test_certs/server.key",
	}
	l, e := tls.Listener(addr)
	assert.NotNil(e)
	assert.Nil(l)

	tls = &web.TLS{
		CertFile: "blah.crt",
		KeyFile:  "blah.key",
	}
	l, e = tls.Listener(addr)
	assert.NotNil(e)
	assert.Nil(l)

	addr = web.NewAddress("gabbar.bad.address.com:8080")
	tls = &web.TLS{
		CertFile: "../test_certs/server.crt",
		KeyFile:  "../test_certs/server.key",
	}
	l, e = tls.Listener(addr)
	assert.NotNil(e)
	assert.Nil(l)
}
