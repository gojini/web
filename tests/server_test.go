package server_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gojini.dev/web"
)

var message = "hello"

var handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(message)); err != nil {
		panic(err)
	}
})

func TestWebServer(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	cfg := &web.Config{
		Address: web.NewAddress("127.0.0.1:9999"),
		TLS:     nil,
	}

	server := web.NewServer(cfg, handler)
	assert.NotNil(server)

	ctx := context.Background()

	go func() {
		assert.NotNil(server.Start(ctx))
	}()

	// Wait for a second for http server to start
	time.Sleep(time.Second)

	// Write client to test the server
	resp, err := http.Get(fmt.Sprintf("http://%s/", cfg.Address))
	assert.Nil(err)
	assert.NotNil(resp)

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal(string(bytes), message)
	assert.Nil(resp.Body.Close())

	assert.Nil(server.Stop(ctx, nil))
}

func TestBadWebServer(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	cfg := &web.Config{
		Address: web.NewAddress("bad.address:9999"),
		TLS:     nil,
	}

	server := web.NewServer(cfg, handler)
	assert.NotNil(server)

	ctx := context.Background()

	assert.NotNil(server.Start(ctx))
}
