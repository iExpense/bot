package imessenger

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
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

	imessenger := &messenger.Messenger{
		AccessToken: accessToken,
		VerifyToken: verifyToken,
	}

	return &Bot{
		imessenger: imessenger,
		quit:       make(chan interface{}, 0),
	}, nil
}

func (b *Bot) String() string {
	return fmt.Sprintf("")
}

func (b *Bot) MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	profile, err := b.imessenger.GetProfile(opts.Sender.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := b.imessenger.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello, %s %s, %s", profile.FirstName, profile.LastName, msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}

func (b *Bot) Serve() {
	listenPort := viper.GetString("port")
	b.imessenger.MessageReceived = b.MessageReceived
	http.HandleFunc("/webhook", b.imessenger.Handler)

	// TODO: use stoppable server
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

func (b *Bot) Stop() {
	close(b.quit)
}
