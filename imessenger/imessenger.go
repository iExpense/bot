package imessenger

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mangalaman93/messenger"
	"github.com/spf13/viper"
)

type Bot struct {
	imessenger *messenger.Messenger
	quit       chan interface{}
}

func NewBot() (*Bot, error) {
	accessToken := viper.GetString("messenger.access_token")
	if accessToken == "" {
		return nil, fmt.Errorf("key %s is empty", accessToken)
	}

	verifyToken := viper.GetString("messenger.verify_token")
	if verifyToken == "" {
		return nil, fmt.Errorf("key %s is empty", verifyToken)
	}

	listenPort := viper.GetString("port")
	if listenPort == "" {
		return nil, fmt.Errorf("key %s is empty", listenPort)
	}

	return &Bot{
		imessenger: messenger.New(messenger.Options{
			Verify:      false,
			VerifyToken: verifyToken,
			Token:       accessToken,
		}),
		quit: make(chan interface{}, 0),
	}, nil
}

func (b *Bot) Serve() {
	listenPort := viper.GetString("port")

	b.imessenger.HandleMessage(func(m messenger.Message, r *messenger.Response) {
		p, err := b.imessenger.ProfileByID(m.Sender.ID)
		if err != nil {
			fmt.Println("[ERROR] Could not get user profile ::", err)
		}
		r.Text(fmt.Sprintf("Hello, %v!", p.FirstName))
	})

	// TODO: use stoppable server
	log.Println("[INFO] Serving messenger bot on port=" + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, b.imessenger.Handler()))
}

func (b *Bot) Stop() {
	close(b.quit)
}
