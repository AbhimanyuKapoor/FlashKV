package commands

import (
	"fmt"
	"sort"

	"github.com/Hari08-dev/flashkv/internal/client"
)

type SMembersCommand struct {
	BaseCommand
}

func (c *SMembersCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: SMEMBERS <key>")
		return
	}

	set, err := c.Store.GetSet(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	members := set.Members()

	if len(members) == 0 {
		fmt.Fprintln(client.Conn, "(empty set)")
		return
	}

	sort.Strings(members)

	for i, member := range members {
		fmt.Fprintf(client.Conn, "%d) %s\n", i+1, member)
	}
}
