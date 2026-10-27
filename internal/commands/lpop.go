package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type LPopCommand struct {
	BaseCommand
}

func (c *LPopCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: LPOP <key>")
		return
	}

	key := args[1]

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	value, ok := list.LPop()

	if !ok {
		fmt.Fprintln(client.Conn, "(nil)")
		return
	}

	fmt.Fprintln(client.Conn, value)
}
