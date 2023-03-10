package server_test

import (
	"testing"

	"github.com/ecshreve/lker/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	srv := server.NewServer()
	assert.NotNil(t, srv)
}
