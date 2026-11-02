package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type LLenCommand struct {
	BaseCommand
}

func (c *LLenCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: LLEN <key>")
		return
	}

	key := args[1]

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	fmt.Fprintln(client.Conn, list.Len())
}
