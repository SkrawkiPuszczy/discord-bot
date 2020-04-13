package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jasonlvhit/gocron"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
	"github.com/skrawkipuszczy/discord-bot/pkg/config"
	"github.com/skrawkipuszczy/discord-bot/pkg/db"
	"github.com/skrawkipuszczy/discord-bot/pkg/discord"
	"github.com/skrawkipuszczy/discord-bot/pkg/geolocation"
	"github.com/skrawkipuszczy/discord-bot/pkg/instagram"
	"github.com/skrawkipuszczy/discord-bot/pkg/meteo"
	"github.com/skrawkipuszczy/discord-bot/pkg/n2yo"
	"github.com/skrawkipuszczy/discord-bot/pkg/scheduler"
)

func main() {
	var err error
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	cacheCl, err := cache.New(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cacheCl.Close()
	dbCl, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbCl.Close()

	d, err := discord.New(cfg.DiscordToken)
	geoLocationClient := geolocation.New(cacheCl)
	if cfg.N2yoEnabled {
		mm := n2yo.New(geoLocationClient, cfg.N2yoAPIKey)
		d.RegisterHandlers(n2yo.NewIssVisualPassHandler(cfg.CommandPrefix, mm))
	}
	defer d.Close()
	if cfg.BurzeDzisAPIKey != "" {
		meteoCl := meteo.New(cfg.BurzeDzisAPIKey, *dbCl.Db, cacheCl.(cache.WetherCitiesCache), cfg.CommandPrefix)
		d.RegisterProvidedHandlers(meteoCl)
	}
	if cfg.InstagramEnabled {
		i := instagram.New(cfg.InstagramUsername, cfg.InstagramPassword, cacheCl.(cache.PhotosCache))
		go instagram.GetHashTagPhotos(i, cfg.InstagramHashtag)
		d.RegisterHandlers(instagram.NewDisplayRandomInstagramPhotoHandler(cfg.CommandPrefix, cacheCl.(cache.PhotosCache)))
	}
	d.RegisterHandlers(discord.NewRandomUserHandler(cfg.CommandPrefix))
	d.RegisterHandlers(discord.NewHelpHandler(cfg.CommandPrefix, d.GetHandlers()))
	d.GetSession().AddHandler(discord.NewAdMessageHandler(cacheCl.(cache.MessagesOnChannelsCache), cfg.AdMessageInterval).AdMessageHandler)
	err = d.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	if cfg.ScheduledMessagesEnabled {
		if cfg.ScheduledConfigFileURL == "" {
			log.Fatalln("Scheduler config is empty")
		}
		log.Println("scheduled messages feature enabled")
		err := scheduler.New(cfg.ScheduledConfigFileURL, d.SendMessage)
		if err != nil {
			log.Println(err)
		}
		<-gocron.Start()
	}
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
