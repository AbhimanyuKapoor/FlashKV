package commands

import (
	"fmt"
	"net"
	"strconv"
)

type ExpireCommand struct {
	BaseCommand
}

func (c *ExpireCommand) Execute(conn net.Conn, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(conn, "ERR invalid arguments")
		return
	}

	seconds, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(conn, "ERR invalid TTL")
		return
	}

	if !c.Store.Expire(args[1], seconds) {
		fmt.Fprintln(conn, "(nil)")
		return
	}

	fmt.Fprintln(conn, "OK")
}
