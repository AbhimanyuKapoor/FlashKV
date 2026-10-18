package commands

import (
	"fmt"
	"net"
	"strings"
)

type KeysCommand struct {
	BaseCommand
}

func (c *KeysCommand) Execute(conn net.Conn, args []string) {

	keys := c.Store.Keys()

	if len(keys) == 0 {
		fmt.Fprintln(conn, "(empty)")
		return
	}

	fmt.Fprintln(conn, strings.Join(keys, " "))
}
