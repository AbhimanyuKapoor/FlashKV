package commands

import (
	"fmt"

	"FlashKV/internal/client"
)

type PingCommand struct{}

func (c *PingCommand) Execute(client *client.Client, args []string) {
	fmt.Fprintln(client.Conn, "PONG")
}
