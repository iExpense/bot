package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iexpense/bot/iconf"
	"github.com/iexpense/bot/ihangout"
	"github.com/thejerf/suture"
)

func main() {
	iconf.Init()

	// register ctrl+c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Println("[INFO] adding signal handler for SIGTERM")

	ihangoutBot, err := ihangout.NewBot()
	if err != nil {
		log.Println("[ERROR] unable to create ihangout bot")
		panic(err)
	}

	supervisor := suture.NewSimple("bot")
	supervisor.Add(ihangoutBot)
	go supervisor.ServeBackground()
	log.Printf("[INFO] running supervisor: %v", supervisor)

	log.Println("[INFO] waiting for ctrl+c signal")
	<-sigs
	supervisor.Stop()
	log.Println("[INFO] exiting bot")
}
