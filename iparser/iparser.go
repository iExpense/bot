package iparser

import (
	"errors"
	"fmt"
	"strings"
)

type CType int

const (
	Expense CType = iota
	Income
	Transfer
	History
	Balance
	Help
)

type Command struct {
	Ctype  CType
	Amount *Money
	Tags   []string
}

var (
	// Error Codes
	ErrorEmptyCmd = errors.New("Empty command")
	ErrorNeedArgs = errors.New("Command needs arguments")

	// constant variables
	historyCommand = &Command{
		Ctype: History,
	}
	balanceCommand = &Command{
		Ctype: Balance,
	}
	helpCommand = &Command{
		Ctype: Help,
	}
)

func (c Command) String() string {
	return fmt.Sprintf("cmd:%d, amount:%s, tags:%s", c.Ctype, c.Amount, c.Tags)
}

func Parse(line string) (*Command, error) {
	line = strings.ToLower(line)
	tokens := strings.Fields(line)

	if len(tokens) <= 0 {
		return nil, ErrorEmptyCmd
	}

	switch cmd := tokens[0]; {
	case strings.HasPrefix("expense", cmd):
		return parseExpense(tokens[1:])
	case strings.HasPrefix("income", cmd):
		return parseIncome(tokens[1:])
	case strings.HasPrefix("transfer", cmd):
		return parseTransfer(tokens[1:])
	case strings.HasPrefix("history", cmd):
		return parseHistory(tokens[1:])
	case strings.HasPrefix("balance", cmd):
		return parseBalance(tokens[1:])
	case strings.EqualFold("help", cmd):
		return parseHelp(tokens[1:])
	default:
		return parseExpense(tokens)
	}
}

func parseExpense(args []string) (*Command, error) {
	if len(args) <= 0 {
		return nil, ErrorNeedArgs
	} else if len(args) <= 1 {
		return nil, ErrorNeedArgs
	}

	amount, err := parseAmount(args[0])
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(args[1:]))
	for _, arg := range args[1:] {
		tags = append(tags, strings.TrimPrefix(arg, "#"))
	}

	return &Command{
		Ctype:  Expense,
		Amount: amount,
		Tags:   tags,
	}, nil
}

func parseIncome(args []string) (*Command, error) {
	if len(args) <= 0 {
		return nil, ErrorNeedArgs
	} else if len(args) <= 1 {
		return nil, ErrorNeedArgs
	}

	amount, err := parseAmount(args[0])
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(args[1:]))
	for _, arg := range args[1:] {
		tags = append(tags, strings.TrimPrefix(arg, "#"))
	}

	return &Command{
		Ctype:  Income,
		Amount: amount,
		Tags:   tags,
	}, nil
}

func parseTransfer(args []string) (*Command, error) {
	if (len(args) < 3) || (args[2] == ">" && len(args) < 4) {
		return nil, ErrorNeedArgs
	}

	amount, err := parseAmount(args[0])
	if err != nil {
		return nil, err
	}

	fromAccount := strings.TrimPrefix(args[1], "#")
	var toAccount string
	if args[2] == ">" {
		toAccount = strings.TrimPrefix(args[3], "#")
	} else {
		toAccount = strings.TrimPrefix(args[2], "#")
	}

	return &Command{
		Ctype:  Transfer,
		Amount: amount,
		Tags:   []string{fromAccount, toAccount},
	}, nil
}

func parseHistory(args []string) (*Command, error) {
	if len(args) == 0 {
		return historyCommand, nil
	}

	timeframe := args[0]
	return &Command{
		Ctype: History,
		Tags:  []string{timeframe},
	}, nil
}

func parseBalance(args []string) (*Command, error) {
	if len(args) == 0 {
		return balanceCommand, nil
	}

	tags := make([]string, 0, len(args))
	for _, arg := range args {
		tags = append(tags, strings.TrimPrefix(arg, "#"))
	}

	return &Command{
		Ctype: Balance,
		Tags:  tags,
	}, nil
}

func parseHelp(args []string) (*Command, error) {
	return helpCommand, nil
}

func parseAmount(amount string) (*Money, error) {
	return NewMoney(amount)
}
