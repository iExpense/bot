package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iexpense/bot/iconf"
	"github.com/iexpense/bot/imessenger"
	"github.com/thejerf/suture"
)

func main() {
	iconf.Init()

	// register ctrl+c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Println("[INFO] adding signal handler for SIGTERM")

	imessengerBot, err := imessenger.NewBot()
	if err != nil {
		log.Println("[ERROR] unable to create imessenger bot")
		panic(err)
	}

	supervisor := suture.NewSimple("bot")
	supervisor.Add(imessengerBot)
	go supervisor.ServeBackground()
	log.Printf("[INFO] running supervisor: %v", supervisor)

	log.Println("[INFO] waiting for ctrl+c signal")
	<-sigs
	supervisor.Stop()
	log.Println("[INFO] exiting bot")
}
