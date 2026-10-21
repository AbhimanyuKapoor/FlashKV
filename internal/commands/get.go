package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type GetCommand struct {
	BaseCommand
}

func (c *GetCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	value, ok := c.Store.Get(args[1])

	if !ok {
		fmt.Fprintln(client.Conn, "(nil)")
		return
	}

	fmt.Fprintln(client.Conn, value)
}
