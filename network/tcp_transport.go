package network

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a peer node in the network.
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

// TCPTransport manages TCP connections between peers.
type TCPTransport struct {
	ListenAddress string
	listener      net.Listener
	handshakeFunc HandshakeFunc
	peers         map[net.Addr]*TCPPeer
	mu            sync.RWMutex
}

// NewTCPTransport creates a new TCP transport.
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddr,
		handshakeFunc: DefaultHandshake,
		peers:         make(map[net.Addr]*TCPPeer),
	}
}

// ListenAndAccept starts the TCP listener and accepts incoming connections.
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	fmt.Println("Listening on", t.ListenAddress)

	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop listens for incoming peer connections.
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}

		peer := &TCPPeer{conn: conn, outbound: false}

		t.mu.Lock()
		t.peers[conn.RemoteAddr()] = peer
		t.mu.Unlock()

		go t.handleConn(peer)
	}
}

// handleConn processes incoming peer connections.
func (t *TCPTransport) handleConn(peer *TCPPeer) {
	defer peer.conn.Close()

	if err := t.handshakeFunc(peer.conn); err != nil {
		fmt.Println("Handshake failed:", err)
		return
	}

	fmt.Printf("New connection established: %s\n", peer.conn.RemoteAddr().String())
}
