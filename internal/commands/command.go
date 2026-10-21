package commands

import (
	"github.com/Hari08-dev/flashkv/internal/client"
	"github.com/Hari08-dev/flashkv/internal/datastore"
)

type Command interface {
	Execute(client *client.Client, args []string)
}

type BaseCommand struct {
	Store *datastore.Store
}
