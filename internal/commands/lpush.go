package commands

import (
	"fmt"
	"strings"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type LPushCommand struct {
	BaseCommand
}

func (c *LPushCommand) Execute(client *client.Client, args []string) {

	if len(args) < 3 {
		fmt.Fprintln(client.Conn, "ERR usage: LPUSH <key> <value>")
		return
	}

	key := args[1]
	value := strings.Join(args[2:], " ")

	list, err := c.Store.GetList(key)

	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	list.LPush(value)

	fmt.Fprintln(client.Conn, "OK")
}
