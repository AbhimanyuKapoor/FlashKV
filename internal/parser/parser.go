package parser

import (
	"fmt"
	"strings"

	"github.com/Hari08-dev/flashkv/internal/client"
	"github.com/Hari08-dev/flashkv/internal/commands"
	"github.com/Hari08-dev/flashkv/internal/datastore"
)

type Parser struct {
	commands map[string]commands.Command
}

func New(store *datastore.Store) *Parser {

	base := commands.BaseCommand{
		Store: store,
	}

	return &Parser{
		commands: map[string]commands.Command{
			"PING": &commands.PingCommand{},

			"SET": &commands.SetCommand{
				BaseCommand: base,
			},

			"GET": &commands.GetCommand{
				BaseCommand: base,
			},

			"DEL": &commands.DelCommand{
				BaseCommand: base,
			},

			"EXISTS": &commands.ExistsCommand{
				BaseCommand: base,
			},

			"EXPIRE": &commands.ExpireCommand{
				BaseCommand: base,
			},

			"TTL": &commands.TTLCommand{
				BaseCommand: base,
			},

			"KEYS": &commands.KeysCommand{
				BaseCommand: base,
			},

			"SAVE": &commands.SaveCommand{
				BaseCommand: base,
			},

			"LOAD": &commands.LoadCommand{
				BaseCommand: base,
			},
		},
	}
}

func (p *Parser) Execute(client *client.Client, command string) {

	fields := strings.Fields(command)

	if len(fields) == 0 {
		fmt.Fprintln(client.Conn, "ERR empty command")
		return
	}

	cmd := strings.ToUpper(fields[0])

	handler, ok := p.commands[cmd]
	if !ok {
		fmt.Fprintln(client.Conn, "ERR unknown command")
		return
	}

	handler.Execute(client, fields)
}
