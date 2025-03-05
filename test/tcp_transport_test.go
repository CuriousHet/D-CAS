package test

import (
	"testing"

	"github.com/CuriousHet/D-CAS/network"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := network.NewTCPTransport(listenAddr)

	assert.Equal(t, tr.ListenAddress, listenAddr)
	assert.Nil(t, tr.ListenAndAccept())
}
