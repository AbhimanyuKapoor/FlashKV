package parser

import (
	"fmt"
	"strings"

	"FlashKV/internal/client"
	"FlashKV/internal/commands"
	"FlashKV/internal/datastore"
	"FlashKV/internal/pubsub"
)

type Parser struct {
	commands map[string]commands.Command
}

func New(store *datastore.Store) *Parser {

	ps := pubsub.New()

	base := commands.BaseCommand{
		Store:  store,
		PubSub: ps,
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
			"PUBLISH": &commands.PublishCommand{
				BaseCommand: base,
			},

			"SUBSCRIBE": &commands.SubscribeCommand{
				BaseCommand: base,
			},

			"LPUSH": &commands.LPushCommand{
				BaseCommand: base,
			},

			"RPUSH": &commands.RPushCommand{
				BaseCommand: base,
			},

			"LPOP": &commands.LPopCommand{
				BaseCommand: base,
			},

			"RPOP": &commands.RPopCommand{
				BaseCommand: base,
			},

			"LLEN": &commands.LLenCommand{
				BaseCommand: base,
			},

			"LRANGE": &commands.LRangeCommand{
				BaseCommand: base,
			},
			"HSET": &commands.HSetCommand{
				BaseCommand: base,
			},
			"HGET": &commands.HGetCommand{
				BaseCommand: base,
			},
			"HDEL": &commands.HDelCommand{
				BaseCommand: base,
			},
			"HEXISTS": &commands.HExistsCommand{
				BaseCommand: base,
			},
			"HGETALL": &commands.HGetAllCommand{
				BaseCommand: base,
			},
			"HLEN": &commands.HLenCommand{
				BaseCommand: base,
			},
			"SADD": &commands.SAddCommand{
				BaseCommand: base,
			},

			"SREM": &commands.SRemCommand{
				BaseCommand: base,
			},

			"SISMEMBER": &commands.SIsMemberCommand{
				BaseCommand: base,
			},

			"SMEMBERS": &commands.SMembersCommand{
				BaseCommand: base,
			},

			"SCARD": &commands.SCardCommand{
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
