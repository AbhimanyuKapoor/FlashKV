package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type LRangeCommand struct {
	BaseCommand
}

func (c *LRangeCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: LRANGE <key>")
		return
	}

	key := args[1]

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	values := list.Range()

	if len(values) == 0 {
		fmt.Fprintln(client.Conn, "(empty list)")
		return
	}

	for i, value := range values {
		fmt.Fprintf(client.Conn, "%d) %s\n", i+1, value)
	}
}
