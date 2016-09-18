package ihangout

import (
	"fmt"
	"time"

	"github.com/mattn/go-xmpp"
	"github.com/spf13/viper"
)

const (
	cHangoutHost    = "hangouts.google.com"
	cHangoutPort    = "1234"
	cPresencePeriod = 100
)

type Bot struct {
	client *xmpp.Client
}

func NewBot() (*Bot, error) {
	username := viper.GetString("hangout.username")
	password := viper.GetString("hangout.password")
	host := fmt.Sprintf("%s:%s", cHangoutHost, cHangoutPort)
	client, err := xmpp.NewClient(host, username, password, true)
	if err != nil {
		return nil, err
	}

	return &Bot{
		client: client,
	}
}

func (b *Bot) String() string {
	return fmt.Printf("client: %v", b.client)
}

func (b *Bot) Serve() {
	ticker := time.NewTicker(cPresencePeriod * time.Second)
	select {
	case <-ticker.C:
		b.client.SendPresence(xmpp.Presence{})
	}
}

func (b *Bot) Stop() {
	b.client.Close()
}
