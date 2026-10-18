package commands

import (
	"net"

	"github.com/Hari08-dev/flashkv/internal/datastore"
	"github.com/Hari08-dev/flashkv/internal/pubsub"
)

type Command interface {
	Execute(conn net.Conn, args []string)
}

type BaseCommand struct {
	Store  *datastore.Store
	PubSub *pubsub.PubSub
}
