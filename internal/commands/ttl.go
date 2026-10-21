package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type TTLCommand struct {
	BaseCommand
}

func (c *TTLCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	fmt.Fprintln(client.Conn, c.Store.TTL(args[1]))
}
