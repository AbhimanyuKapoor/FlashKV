package commands

import (
	"fmt"
	"net"

	"github.com/Hari08-dev/flashkv/internal/persistence"
)

type LoadCommand struct {
	BaseCommand
}

func (c *LoadCommand) Execute(conn net.Conn, args []string) {

	if err := persistence.Load("snapshots/dump.json", c.Store); err != nil {
		fmt.Fprintln(conn, err.Error())
		return
	}

	fmt.Fprintln(conn, "OK")
}
