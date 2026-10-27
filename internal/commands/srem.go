package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type SRemCommand struct {
	BaseCommand
}

func (c *SRemCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR usage: SREM <key> <member>")
		return
	}

	set, err := c.Store.GetSet(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	if set.Remove(args[2]) {
		fmt.Fprintln(client.Conn, "1")
	} else {
		fmt.Fprintln(client.Conn, "0")
	}
}
