package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jasonlvhit/gocron"
	"github.com/kelseyhightower/envconfig"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
	"github.com/skrawkipuszczy/discord-bot/pkg/config"
	"github.com/skrawkipuszczy/discord-bot/pkg/discord"
	"github.com/skrawkipuszczy/discord-bot/pkg/geolocation"
	"github.com/skrawkipuszczy/discord-bot/pkg/instagram"
	"github.com/skrawkipuszczy/discord-bot/pkg/n2yo"
	"github.com/skrawkipuszczy/discord-bot/pkg/scheduler"
)

func main() {
	var err error
	var c config.EnvConfig
	err = envconfig.Process("bot", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	var cacheCl cache.Cache
	if c.CacheType == "redis" {
		cacheCl, err = cache.NewRedisClient(c.RedisUrl)
	} else {
		cacheCl, err = cache.NewMemoryCache()
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cacheCl.Close()
	d, err := discord.New(c.DiscordToken)
	if c.N2yoEnabled {
		geoLocationClient := geolocation.New(cacheCl)
		mm := n2yo.New(geoLocationClient, c.N2yoApiKey)
		d.RegisterHandlers(n2yo.NewIssVisualPassHandler(c.CommandPrefix, mm))
	}
	defer d.Close()
	if c.InstagramEnabled {
		i := instagram.New(c.InstagramUsername, c.InstagramPassword, cacheCl.(cache.PhotosCache))
		go i.GetHashTagPhotos(c.InstagramHashtag)
		d.RegisterHandlers(instagram.NewDisplayRandomInstagramPhotoHandler(c.CommandPrefix, cacheCl.(cache.PhotosCache)))
	}
	d.RegisterHandlers(discord.NewHelpHandler(c.CommandPrefix, d.GetHandlers()))
	err = d.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	if c.ScheduledMessagesEnabled {
		if c.ScheduledConfigFileUrl == "" {
			log.Fatalln("Scheduler config is empty")
		}
		log.Println("scheduled messages feature enabled")
		err := scheduler.New(c.ScheduledConfigFileUrl, d.SendMessage)
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
