package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type DelCommand struct {
	BaseCommand
}

func (c *DelCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	if c.Store.Delete(args[1]) {
		fmt.Fprintln(client.Conn, "OK")
		return
	}

	fmt.Fprintln(client.Conn, "(nil)")
}
