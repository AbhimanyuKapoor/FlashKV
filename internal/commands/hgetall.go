package commands

import (
	"fmt"
	"sort"

	"FlashKV/internal/client"
)

type HGetAllCommand struct {
	BaseCommand
}

func (c *HGetAllCommand) Execute(client *client.Client, args []string) {

	if len(args) != 2 {
		fmt.Fprintln(client.Conn, "ERR usage: HGETALL <key>")
		return
	}

	hash, err := c.Store.GetHash(args[1])
	if err != nil {
		fmt.Fprintln(client.Conn, err)
		return
	}

	data := hash.GetAll()

	if len(data) == 0 {
		fmt.Fprintln(client.Conn, "(empty hash)")
		return
	}

	fields := make([]string, 0, len(data))
	for k := range data {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	i := 1
	for _, field := range fields {
		fmt.Fprintf(client.Conn, "%d) %s : %s\n", i, field, data[field])
		i++
	}
}
