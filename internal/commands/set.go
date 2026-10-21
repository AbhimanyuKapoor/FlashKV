package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type SetCommand struct {
	BaseCommand
}

func (c *SetCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	c.Store.Set(args[1], args[2])

	fmt.Fprintln(client.Conn, "OK")
}
