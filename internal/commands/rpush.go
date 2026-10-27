package commands

import (
	"fmt"
	"strings"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type RPushCommand struct {
	BaseCommand
}

func (c *RPushCommand) Execute(client *client.Client, args []string) {

	if len(args) < 3 {
		fmt.Fprintln(client.Conn, "ERR usage: RPUSH <key> <value>")
		return
	}

	key := args[1]
	value := strings.Join(args[2:], " ")

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	list.RPush(value)

	fmt.Fprintln(client.Conn, "OK")
}
