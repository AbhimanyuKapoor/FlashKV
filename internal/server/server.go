package server

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"

	"github.com/Hari08-dev/flashkv/internal/datastore"
	"github.com/Hari08-dev/flashkv/internal/parser"
	"github.com/Hari08-dev/flashkv/internal/ttl"
)

type Server struct {
	address string
	store   *datastore.Store
	parser  *parser.Parser
}

func NewServer(address string) *Server {

	store := datastore.NewStore()

	cleaner := ttl.New(store, time.Second)

	go cleaner.Start()

	return &Server{
		address: address,
		store:   store,
		parser:  parser.New(store),
	}
}

func (s *Server) Start() error {

	// Start TCP listener
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Listening on %s\n", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("New client connected: %s\n", conn.RemoteAddr())

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		// Show prompt
		if _, err := conn.Write([]byte("flashkv> ")); err != nil {
			return
		}

		// Read command
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client disconnected: %s\n", conn.RemoteAddr())
			return
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		log.Printf("[%s] %s\n", conn.RemoteAddr(), command)

		// Execute command
		s.parser.Execute(conn, command)
	}
}
