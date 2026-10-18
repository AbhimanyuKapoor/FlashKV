package commands

import (
	"fmt"
	"net"
)

type GetCommand struct {
	BaseCommand
}

func (c *GetCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	value, ok := c.Store.Get(args[1])

	if !ok {
		fmt.Fprintln(conn, "(nil)")
		return
	}

	fmt.Fprintln(conn, value)
}
