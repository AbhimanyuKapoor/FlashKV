package commands

import (
	"fmt"
	"strings"

	"FlashKV/internal/client"
)

type HSetCommand struct {
	BaseCommand
}

func (c *HSetCommand) Execute(client *client.Client, args []string) {

	if len(args) < 4 {
		fmt.Fprintln(client.Conn, "ERR usage: HSET <key> <field> <value>")
		return
	}

	hash, err := c.Store.GetHash(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	field := args[2]
	value := strings.Join(args[3:], " ")

	hash.Set(field, value)

	fmt.Fprintln(client.Conn, "OK")
}
