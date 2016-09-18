package iparser

import (
	"errors"
	"strings"
)

type EType int

const (
	Balance EType = iota
	Expense
	History
	Income
	Transfer
)

// Error Codes
var (
	ErrorEmptyCmd   = errors.New("iexpense: empty command")
	ErrorInvalidCmd = errors.New("iexpense: invalid command")
)

type Command struct {
	etype EType
	args  []string
}

func Parse(line string) (*Command, error) {
	line = strings.ToLower(line)
	tokens := strings.Fields(line)

	if len(tokens) <= 0 {
		return nil, ErrorEmptyCmd
	}

	switch cmd := tokens[0]; cmd {
	case strings.HasPrefix("transfer", cmd):
		return parseTransfer(tokens[1:])
	case strings.HasPrefix("income", cmd):
		return parseIncome(tokens[1:])
	case strings.HasPrefix("history", cmd):
		return parseHistory(tokens[1:])
	case strings.HasPrefix("balance", cmd):
		return parseBalance(tokens[1:])
	default:
		return parseExpense(tokens)
	}
}

func parseTransfer(args []string) (*Command, error) {
	return nil, nil
}

func parseIncome(args []string) (*Command, error) {
	return nil, nil
}

func parseHistory(args []string) (*Command, error) {
	return nil, nil
}

func parseBalance(args []string) (*Command, error) {
	return nil, nil
}

func parseExpense(args []string) (*Command, error) {
	return nil, nil
}
