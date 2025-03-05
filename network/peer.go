package network

import "net"

// Peer represents a node in the distributed network.
type Peer interface {
	Address() net.Addr
	SendMessage(data []byte) error
}
