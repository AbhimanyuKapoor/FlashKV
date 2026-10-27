package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type HGetCommand struct {
	BaseCommand
}

func (c *HGetCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR usage: HGET <key> <field>")
		return
	}

	hash, err := c.Store.GetHash(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	value, ok := hash.Get(args[2])
	if !ok {
		fmt.Fprintln(client.Conn, "(nil)")
		return
	}

	fmt.Fprintln(client.Conn, value)
}
