package iexpense

import (
	"github.com/iexpense/bot/iparser"
	"github.com/mangalaman93/messenger"
)

type EType int

const (
	Expense EType = iota
	Income
	Transfer
)

type IExpense struct {
	Etype  EType
	Amount Money
	Tags   []string
}

func (b *Bot) HandleReceivedMessage(m messenger.Message, r *messenger.Response) {
	var reply string

	cmd, err := iparser.Parse(m.Text)
	if err != nil {
		reply = err.Error()
	} else {
		reply = cmd.String()
	}

	r.Text(reply)
}
