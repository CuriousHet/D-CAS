package network

import (
	"fmt"
	"net"
)

// HandshakeFunc represents a function for peer handshaking.
type HandshakeFunc func(conn net.Conn) error

// Default handshake function that simply accepts the connection.
func DefaultHandshake(conn net.Conn) error {
	fmt.Printf("Handshake with %s\n", conn.RemoteAddr().String())
	return nil
}
