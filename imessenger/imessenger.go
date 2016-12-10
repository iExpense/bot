package imessenger

import (
	"fmt"
	"log"
	"time"
)

const (
	cPresencePeriod = 100
)

type Bot struct {
	quit chan interface{}
}

func NewBot() (*Bot, error) {
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
