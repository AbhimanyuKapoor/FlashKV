package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type PingCommand struct{}

func (c *PingCommand) Execute(client *client.Client, args []string) {
	fmt.Fprintln(client.Conn, "PONG")
}
