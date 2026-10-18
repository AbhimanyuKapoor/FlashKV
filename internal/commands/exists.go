package commands

import (
	"fmt"
	"net"
)

type ExistsCommand struct {
	BaseCommand
}

func (c *ExistsCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	if c.Store.Exists(args[1]) {
		fmt.Fprintln(conn, "1")
		return
	}

	fmt.Fprintln(conn, "0")
}
