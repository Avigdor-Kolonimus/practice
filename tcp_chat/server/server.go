package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"tcp_chat/message"
	"tcp_chat/peer"
)

const (
	messageLimit = 100 // limit of messages in the channel
	bufferSize = 2048 // size of the buffer
)

type Server struct {
	Address  string      // host:port
	Listener net.Listener // net.Listener is an interface that represents a listener for incoming connections
	messagesChan  chan message.Message // channel for messages
	clients      map[net.Conn]*peer.Peer // map of clients
	deadClients  []net.Conn // slice of dead clients
	mu sync.RWMutex // mutex for the clients map
}


func NewServer(address string) *Server {
	return &Server{
		Address: address,
		messagesChan: make(chan message.Message, messageLimit),
		clients: make(map[net.Conn]*peer.Peer),
		deadClients: make([]net.Conn, 0),
		mu: sync.RWMutex{},
	}
}

// ------------------------------------------------------------ Start ------------------------------------------------------------

func (s *Server) Start() error {
	var err error

	s.Listener, err = net.Listen("tcp", s.Address)
	if err != nil {
		log.Printf("Error accept: %v", err)

		return err
	}

	log.Printf("Server started")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	
	go s.acceptLoop(ctx)
	go s.Broadcast(ctx)
	
	defer func() {
		cancel()
		s.Listener.Close()
		close(s.messagesChan)
	}()
	s.Stop(ctx)

	return nil
}

// ------------------------------------------------------------ Stop ------------------------------------------------------------

func (s *Server) Stop(ctx context.Context) {
	<-ctx.Done()
	for _, client := range s.clients {
		s.unregisterPeer(client.Conn)
	}
}

// ------------------------------------------------------------ Broadcast ------------------------------------------------------------

func (s *Server) Broadcast(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-s.messagesChan:
			s.mu.RLock()
			message := fmt.Sprintf("%s: %s\n", msg.Author, msg.Text)
			for _, client := range s.clients {
				s.writeInConnection(client.Conn, message)
			}
			
			s.mu.RUnlock()
			for _, conn := range s.deadClients {
				s.unregisterPeer(conn)
			}
			s.deadClients = s.deadClients[:0]
		}
	}
}

// ------------------------------------------------------------ Utils ------------------------------------------------------------

func (s *Server) acceptLoop(ctx context.Context) {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				log.Println("Accept loop stopped")

				return

			default:
				log.Printf("Failed accept client: %v", err)
			}
		}

		log.Printf("Welcome, %s", conn.RemoteAddr().String())
		
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	s.registerPeer(conn)
	buf := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				s.unregisterPeer(conn)
				return // client closed connection normally
			}

			log.Printf("Connection error: %v", err)
			
			return
		}

		msg := &message.Message{
			Author: conn.RemoteAddr().String(),
			Text:   string(buf[:n]),
		}
		s.messagesChan <- *msg
	}
}

func (s *Server) writeInConnection(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Printf("Failed write message: %v", err)
		s.deadClients = append(s.deadClients, conn)
	}
}

func (s *Server) registerPeer(conn net.Conn) {
	peer := &peer.Peer{
		Conn:        conn,
		ConnectedAt: time.Now(),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[conn] = peer
}

func (s *Server) unregisterPeer(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	conn.Close()
	delete(s.clients, conn)

	log.Printf("Client disconnected: %s", conn.RemoteAddr().String())
}