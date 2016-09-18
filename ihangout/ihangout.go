package ihangout

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	cHangoutHost    = "hangouts.google.com"
	cHangoutPort    = "1234"
	cPresencePeriod = 100
)

type Bot struct {
	quit chan interface{}
}

func NewBot() (*Bot, error) {
	username := viper.GetString("hangout.username")
	password := viper.GetString("hangout.password")
	host := fmt.Sprintf("%s:%s", cHangoutHost, cHangoutPort)
	log.Printf("user: %s, pass: %s, host:%s", username, password, host)

	return &Bot{
		quit: make(chan interface{}, 0),
	}, nil
}

func (b *Bot) String() string {
	return fmt.Sprintf("")
}

func (b *Bot) Serve() {
	ticker := time.NewTicker(cPresencePeriod * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("send presence here")
		case <-b.quit:
			return
		}
	}
}

func (b *Bot) Stop() {
	close(b.quit)
}
