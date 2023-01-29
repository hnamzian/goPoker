package p2p

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

type ServerConfig struct {
	ListenAddr string
	Version    string
}

type Peer struct {
	conn net.Conn
}

type Message struct {
	from    net.Addr
	payload io.Reader
}

type Server struct {
	ServerConfig

	listener net.Listener

	handler Handler

	peers map[net.Addr]*Peer

	addPeer chan *Peer
	delPeer chan *Peer
	msgchan chan *Message
}

func New(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,

		handler: &Defaulthandler{},

		peers: make(map[net.Addr]*Peer),

		addPeer: make(chan *Peer),
		delPeer: make(chan *Peer),
		msgchan: make(chan *Message),
	}
}

func (s *Server) Start() error {
	go s.loop()

	if err := s.listen(); err != nil {
		return err
	}

	s.acceptLoop()

	return nil
}

func (s *Server) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("connect remote failed: %s", err)
		return fmt.Errorf("connect remote failed: %s", err)
	}

	s.addPeer <- &Peer{conn}

	return nil
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ServerConfig.ListenAddr)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("start server failed: %s", err))
	}

	fmt.Printf("Game Server is Running on Port %s\n", s.ServerConfig.ListenAddr)

	s.listener = ln

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(fmt.Sprintf("accept connection failed: %s", err))
		}

		p := &Peer{conn}

		// add peer to Peers List
		s.addPeer <- p

		// handle read from peer
		go s.readLoop(p)
	}
}

func (s *Server) readLoop(p *Peer) {
	buff := make([]byte, 1024)

	for {
		n, err := p.conn.Read(buff)
		if err != nil {
			break
		}

		s.msgchan <- &Message{
			from:    p.conn.RemoteAddr(),
			payload: bytes.NewReader(buff[:n]),
		}
	}

	// delete peer from Peers List
	s.delPeer <- p
}

func (s *Server) handleAddPeer(p *Peer) {
	fmt.Printf("New Player Connected: %s\n", p.conn.RemoteAddr())
	s.peers[p.conn.RemoteAddr()] = p
	p.conn.Write([]byte(s.ServerConfig.Version))
}

func (s *Server) handleDeletepeer(p *Peer) {
	fmt.Printf("Player Disconnected: %s\n", p.conn.RemoteAddr())
	delete(s.peers, p.conn.RemoteAddr())
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeer:
			s.handleAddPeer(peer)

		case peer := <-s.delPeer:
			s.handleDeletepeer(peer)

		case msg := <-s.msgchan:
			s.handler.HandleMessage(msg)
		}
	}
}
