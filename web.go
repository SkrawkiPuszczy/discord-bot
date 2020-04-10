package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/skrawkipuszczy/discord-bot/pkg/config"
	"github.com/skrawkipuszczy/discord-bot/pkg/http"
)

func main() {
	var err error
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	svr := http.New(cfg)

	defer svr.Close()
	log.Println("Web app is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		<-sc
		log.Println("receive interrupt signal")
		if err := svr.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()
	err = svr.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
