package instagram

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

type instagramClient struct {
	cl    *goinsta.Instagram
	cache cache.PhotosCache
}

func New(username, password string, cache cache.PhotosCache) *instagramClient {
	insta := goinsta.New(username, password)
	return &instagramClient{cl: insta, cache: cache}
}

func GetHashTagPhotos(i *instagramClient, name string) error {
	for {
		if err := i.cl.Login(); err != nil {
			log.Println(err)
			return err
		}
		defer i.cl.Logout()
		feedTag, err := i.cl.Feed.Tags(name)
		if err != nil {
			log.Println(err)
			return err
		}
		for feedTag.Next() {
			for _, item := range feedTag.Images {
				data, err := json.Marshal(item)
				if err != nil {
					log.Println(err)
					return err
				}
				err = i.cache.SetPhoto(item.ID, string(data))
				if err != nil {
					log.Println(err)
					return err
				}
			}
			min := 5
			max := 120
			sleepTime := rand.Intn(max-min) + min
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
		time.Sleep(time.Duration(24) * time.Hour)
	}

	return nil
}
