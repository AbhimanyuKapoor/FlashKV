package commands

import (
	"fmt"
	"strings"

	"FlashKV/internal/client"
)

type PublishCommand struct {
	BaseCommand
}

func (c *PublishCommand) Execute(client *client.Client, args []string) {

	if len(args) < 3 {
		fmt.Fprintln(client.Conn, "ERR invalid arguments")
		return
	}

	channel := args[1]

	message := strings.Join(args[2:], " ")

	count := c.PubSub.Publish(channel, message)

	fmt.Fprintf(
		client.Conn,
		"Delivered to %d subscriber(s)\n",
		count,
	)
}
