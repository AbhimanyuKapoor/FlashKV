package commands

import (
	"fmt"
	"net"
)

type PingCommand struct{}

func (c *PingCommand) Execute(conn net.Conn, args []string) {
	fmt.Fprintln(conn, "PONG")
}
