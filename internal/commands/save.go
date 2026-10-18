package commands

import (
	"fmt"
	"net"

	"github.com/Hari08-dev/flashkv/internal/persistence"
)

type SaveCommand struct {
	BaseCommand
}

func (c *SaveCommand) Execute(conn net.Conn, args []string) {

	if err := persistence.Save("snapshots/dump.json", c.Store); err != nil {
		fmt.Fprintln(conn, err.Error())
		return
	}

	fmt.Fprintln(conn, "OK")
}
