package ibot

import (
	"github.com/iexpense/bot/iparser"
	"github.com/paked/messenger"
)

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
