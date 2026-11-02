package commands

import (
	"FlashKV/internal/client"
	"FlashKV/internal/datastore"
	"FlashKV/internal/pubsub"
)

type Command interface {
	Execute(client *client.Client, args []string)
}

type BaseCommand struct {
	Store  *datastore.Store
	PubSub *pubsub.PubSub
}
