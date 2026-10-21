package client

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

type Client struct {
	ID   string
	Conn net.Conn

	mu sync.RWMutex

	Subscriptions map[string]chan string
}

func New(conn net.Conn) *Client {

	return &Client{
		ID:            uuid.NewString(),
		Conn:          conn,
		Subscriptions: make(map[string]chan string),
	}
}
