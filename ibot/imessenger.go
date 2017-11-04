package ibot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iexpense/bot/fireclient"
	"github.com/iexpense/bot/iparser"
	"github.com/paked/messenger"
	"github.com/spf13/viper"
)

type Bot struct {
	imessenger *messenger.Messenger
	fc         *fireclient.Fireclient
	quit       chan interface{}
}

func NewBot(fc *fireclient.Fireclient) (*Bot, error) {
	accessToken := viper.GetString("messenger.access_token")
	if accessToken == "" {
		return nil, fmt.Errorf("key messenger.access_token is empty")
	}

	verifyToken := viper.GetString("messenger.verify_token")
	if verifyToken == "" {
		return nil, fmt.Errorf("key messenger.verify_token is empty")
	}

	listenPort := viper.GetString("port")
	if listenPort == "" {
		return nil, fmt.Errorf("key port is empty")
	}

	return &Bot{
		imessenger: messenger.New(messenger.Options{
			Verify:      false,
			VerifyToken: verifyToken,
			Token:       accessToken,
		}),
		fc:   fc,
		quit: make(chan interface{}, 0),
	}, nil
}

func (b *Bot) Serve() {
	listenPort := viper.GetString("port")
	b.imessenger.HandleMessage(b.HandleReceivedMessage)

	// TODO: use stoppable server
	log.Println("[INFO] Serving messenger bot on port=" + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, b.imessenger.Handler()))
}

func (b *Bot) Stop() {
	close(b.quit)
}

func (b *Bot) HandleReceivedMessage(m messenger.Message, r *messenger.Response) {
	var response string
	var errRes error

	if cmd, err := iparser.Parse(m.Text); err != nil {
		errRes = err
	} else {
		switch cmd.Ctype {
		case iparser.Expense:
			response, errRes = b.fc.HandleExpenseCommand(cmd)
		default:
			log.Printf("[ERROR] Command: %+v\n", cmd)
			panic("should not reach here!")
		}
	}

	var reply string
	if errRes != nil {
		reply = errRes.Error()
	} else {
		reply = response
	}
	r.Text(reply)
}
