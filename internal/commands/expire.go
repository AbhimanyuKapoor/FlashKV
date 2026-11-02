package commands

import (
	"fmt"
	"strconv"

	"FlashKV/internal/client"
)

type ExpireCommand struct {
	BaseCommand
}

func (c *ExpireCommand) Execute(client *client.Client, args []string) {

	if len(args) != 3 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	seconds, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(client.Conn, "ERR invalid TTL")
		return
	}

	if !c.Store.Expire(args[1], seconds) {
		fmt.Fprintln(client.Conn, "(nil)")
		return
	}

	fmt.Fprintln(client.Conn, "OK")
}
