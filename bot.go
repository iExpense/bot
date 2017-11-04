package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iexpense/bot/fireclient"
	"github.com/iexpense/bot/ibot"
	"github.com/iexpense/bot/iconf"
	"github.com/thejerf/suture"
)

func main() {
	iconf.Init()

	// register ctrl+c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Println("[INFO] adding signal handler for SIGTERM")

	fc, err := fireclient.NewFireclient()
	if err != nil {
		log.Printf("[ERROR] unable to create fireclient :: %v\n", err)
		panic(err)
	}

	imessengerBot, err := ibot.NewBot(fc)
	if err != nil {
		log.Printf("[ERROR] unable to create imessenger bot :: %v\n", err)
		panic(err)
	}

	supervisor := suture.NewSimple("bot")
	supervisor.Add(imessengerBot)
	go supervisor.ServeBackground()
	log.Printf("[INFO] running supervisor: %v\n", supervisor)

	log.Println("[INFO] waiting for ctrl+c signal")
	<-sigs
	supervisor.Stop()
	log.Println("[INFO] exiting bot")
}
