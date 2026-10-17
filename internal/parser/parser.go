package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Hari08-dev/flashkv/internal/datastore"
	"github.com/Hari08-dev/flashkv/internal/persistence"
)

type Parser struct {
	store    *datastore.Store
	commands map[string]func([]string) string
}

func New(store *datastore.Store) *Parser {

	p := &Parser{
		store: store,
	}

	p.commands = map[string]func([]string) string{
		"PING":   p.handlePing,
		"SET":    p.handleSet,
		"GET":    p.handleGet,
		"DEL":    p.handleDel,
		"EXISTS": p.handleExists,
		"EXPIRE": p.handleExpire,
		"TTL":    p.handleTTL,
		"KEYS":   p.handleKeys,
		"SAVE":   p.handleSave,
		"LOAD":   p.handleLoad,
	}

	return p
}

func (p *Parser) Execute(command string) string {

	fields := strings.Fields(command)

	if len(fields) == 0 {
		return "ERR empty command"
	}

	cmd := strings.ToUpper(fields[0])

	handler, ok := p.commands[cmd]

	if !ok {
		return "ERR unknown command"
	}

	return handler(fields)
}

func (p *Parser) handlePing(args []string) string {
	return "PONG"
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

func (p *Parser) handleExpire(args []string) string {

	if err := validateArgs(args, 3); err != nil {
		return err.Error()
	}

	seconds, err := strconv.Atoi(args[2])
	if err != nil {
		return "ERR invalid TTL"
	}

	if !p.store.Expire(args[1], seconds) {
		return "(nil)"
	}

	return "OK"
}

func (p *Parser) handleTTL(args []string) string {

	if err := validateArgs(args, 2); err != nil {
		return err.Error()
	}

	return strconv.Itoa(p.store.TTL(args[1]))
}

func (p *Parser) handleKeys(args []string) string {

	keys := p.store.Keys()

	if len(keys) == 0 {
		return "(empty)"
	}

	return strings.Join(keys, " ")
}

func (p *Parser) handleSave(args []string) string {

	err := persistence.Save("snapshots/dump.json", p.store)

	if err != nil {
		return err.Error()
	}

	return "OK"
}

func (p *Parser) handleLoad(args []string) string {

	err := persistence.Load("snapshots/dump.json", p.store)

	if err != nil {
		return err.Error()
	}

	return "OK"
}
