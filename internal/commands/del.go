package commands

import (
	"fmt"
	"net"
)

type DelCommand struct {
	BaseCommand
}

func (c *DelCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	if c.Store.Delete(args[1]) {
		fmt.Fprintln(conn, "OK")
		return
	}

	fmt.Fprintln(conn, "(nil)")
}
