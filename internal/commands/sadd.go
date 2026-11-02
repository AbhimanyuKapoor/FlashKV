package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type SAddCommand struct {
	BaseCommand
}

func (c *SAddCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR usage: SADD <key> <member>")
		return
	}

	set, err := c.Store.GetSet(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	if set.Add(args[2]) {
		fmt.Fprintln(client.Conn, "1")
	} else {
		fmt.Fprintln(client.Conn, "0")
	}
}
