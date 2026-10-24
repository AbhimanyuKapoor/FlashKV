package commands

import (
	"fmt"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type SubscribeCommand struct {
	BaseCommand
}

func (c *SubscribeCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	channel := args[1]

	c.PubSub.Subscribe(channel, client)

	fmt.Fprintf(
		client.Conn,
		"Subscribed to %s\n",
		channel,
	)
}
