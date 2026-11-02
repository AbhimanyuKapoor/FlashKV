package server

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"

	"FlashKV/internal/client"
	"FlashKV/internal/datastore"
	"FlashKV/internal/parser"
	"FlashKV/internal/ttl"
)

type Server struct {
	address string
	store   *datastore.Store
	parser  *parser.Parser
	manager *client.Manager
}

func NewServer(address string) *Server {

	store := datastore.NewStore()

	cleaner := ttl.New(store, time.Second)

	manager := client.NewManager()

	go cleaner.Start()

	return &Server{
		address: address,
		store:   store,
		parser:  parser.New(store),
		manager: manager,
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

	client := client.New(conn)

	s.manager.Add(client)

	defer func() {
		s.manager.Remove(client.ID)
		conn.Close()
		log.Printf("Client disconnected: %s\n", client.ID)
	}()

	reader := bufio.NewReader(conn)

	for {

		// Read command
		command, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		log.Printf("[%s] %s\n", client.ID, command)

		// Execute command
		s.parser.Execute(client, command)
	}
}
