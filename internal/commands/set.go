package commands

import (
	"fmt"
	"net"
)

type SetCommand struct {
	BaseCommand
}

func (c *SetCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	c.Store.Set(args[1], args[2])

	fmt.Fprintln(conn, "OK")
}
