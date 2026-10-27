package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type SCardCommand struct {
	BaseCommand
}

func (c *SCardCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: SCARD <key>")
		return
	}

	set, err := c.Store.GetSet(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	fmt.Fprintln(client.Conn, set.Len())
}
