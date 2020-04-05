package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/skrawkipuszczy/discord-bot/pkg/discord"
	"github.com/skrawkipuszczy/discord-bot/pkg/instagram"
	"github.com/skrawkipuszczy/discord-bot/pkg/n2yo"
)

type Config struct {
	DiscordToken      string `required:"true" split_words:"true"`
	CommandPrefix     string `required:"true" split_words:"true"`
	InstagramEnabled  bool   `split_words:"true" default: "false"`
	InstagramUsername string `split_words:"true"`
	InstagramPassword string `split_words:"true"`
	InstagramHashtag  string `split_words:"true"`
	N2yoEnabled       bool   `split_words:"true" default: "false"`
	N2yoApiKey        string `split_words:"true"`
}

func main() {
	var c Config
	err := envconfig.Process("bot", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	d, err := discord.New(c.DiscordToken)
	if c.N2yoEnabled {
		mm := n2yo.New(c.N2yoApiKey)
		d.RegisterHandlers(n2yo.NewIssVisualPassHandler(c.CommandPrefix, mm))
	}
	if c.InstagramEnabled {
		i := instagram.New(c.InstagramUsername, c.InstagramPassword)
		go i.GetHashTagPhotos(c.InstagramHashtag)
		d.RegisterHandlers(instagram.NewDisplayRandomInstagramPhotoHandler(c.CommandPrefix, i.GetPhotos()))
	}
	d.RegisterHandlers(discord.NewHelpHandler(c.CommandPrefix, d.GetHandlers()))
	err = d.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	d.Close()
}
