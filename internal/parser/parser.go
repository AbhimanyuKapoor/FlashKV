package parser

import (
	"errors"
	"strings"

	"github.com/Hari08-dev/flashkv/internal/datastore"
)

type Parser struct {
	store *datastore.Store
}

func New(store *datastore.Store) *Parser {
	return &Parser{
		store: store,
	}
}

func (p *Parser) Execute(command string) string {

	fields := strings.Fields(command)

	if len(fields) == 0 {
		return "ERR empty command"
	}

	switch strings.ToUpper(fields[0]) {

	case "SET":
		return p.handleSet(fields)

	case "GET":
		return p.handleGet(fields)

	case "DEL":
		return p.handleDel(fields)

	case "EXISTS":
		return p.handleExists(fields)

	default:
		return "ERR unknown command"
	}
}

func (p *Parser) handleSet(args []string) string {

	if err := validateArgs(args, 3); err != nil {
		return err.Error()
	}

	p.store.Set(args[1], args[2])

	return "OK"
}

func (p *Parser) handleGet(args []string) string {

	if err := validateArgs(args, 2); err != nil {
		return err.Error()
	}

	value, ok := p.store.Get(args[1])

	if !ok {
		return "(nil)"
	}

	return value
}

func (p *Parser) handleDel(args []string) string {

	if err := validateArgs(args, 2); err != nil {
		return err.Error()
	}

	if p.store.Delete(args[1]) {
		return "OK"
	}

	return "(nil)"
}

func (p *Parser) handleExists(args []string) string {

	if err := validateArgs(args, 2); err != nil {
		return err.Error()
	}

	if p.store.Exists(args[1]) {
		return "1"
	}

	return "0"
}

func validateArgs(args []string, expected int) error {

	if len(args) != expected {
		return errors.New("ERR invalid arguments")
	}

	return nil
}
