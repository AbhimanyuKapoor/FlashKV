package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type HLenCommand struct {
	BaseCommand
}

func (c *HLenCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: HLEN <key>")
		return
	}

	hash, err := c.Store.GetHash(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	fmt.Fprintln(client.Conn, hash.Len())
}
