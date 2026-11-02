package commands

import (
	"fmt"

	"FlashKV/internal/client"
	"FlashKV/internal/persistence"
)

type SaveCommand struct {
	BaseCommand
}

func (c *SaveCommand) Execute(client *client.Client, args []string) {

	if err := persistence.Save("snapshots/dump.json", c.Store); err != nil {
		fmt.Fprintln(client.Conn, err.Error())
		return
	}

	fmt.Fprintln(client.Conn, "OK")
}
