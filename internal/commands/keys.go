package commands

import (
	"fmt"
	"strings"

	"FlashKV/internal/client"
)

type KeysCommand struct {
	BaseCommand
}

func (c *KeysCommand) Execute(client *client.Client, args []string) {

	keys := c.Store.Keys()

	if len(keys) == 0 {
		fmt.Fprintln(client.Conn, "(empty)")
		return
	}

	fmt.Fprintln(client.Conn, strings.Join(keys, " "))
}
