package commands

import (
	"fmt"
	"net"
)

type TTLCommand struct {
	BaseCommand
}

func (c *TTLCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	fmt.Fprintln(conn, c.Store.TTL(args[1]))
}
