package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type RPopCommand struct {
	BaseCommand
}

func (c *RPopCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: RPOP <key>")
		return
	}

	key := args[1]

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	value, ok := list.RPop()

	if !ok {
		fmt.Fprintln(client.Conn, "(nil)")
		return
	}

	fmt.Fprintln(client.Conn, value)
}
