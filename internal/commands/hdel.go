package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type HDelCommand struct {
	BaseCommand
}

func (c *HDelCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR usage: HDEL <key> <field>")
		return
	}

	hash, err := c.Store.GetHash(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	if hash.Delete(args[2]) {
		fmt.Fprintln(client.Conn, "1")
	} else {
		fmt.Fprintln(client.Conn, "0")
	}
}
