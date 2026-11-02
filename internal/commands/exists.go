package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type ExistsCommand struct {
	BaseCommand
}

func (c *ExistsCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	if c.Store.Exists(args[1]) {
		fmt.Fprintln(client.Conn, "1")
		return
	}

	fmt.Fprintln(client.Conn, "0")
}
