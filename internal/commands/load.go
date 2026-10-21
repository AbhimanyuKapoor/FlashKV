package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
	"github.com/Hari08-dev/flashkv/internal/persistence"
)

type LoadCommand struct {
	BaseCommand
}

func (c *LoadCommand) Execute(client *client.Client, args []string) {

	if err := persistence.Load("snapshots/dump.json", c.Store); err != nil {
		fmt.Fprintln(client.Conn, err.Error())
		return
	}

	fmt.Fprintln(client.Conn, "OK")
}
