package instagram

import (
	"log"
	"math/rand"
	"time"

	"github.com/ahmdrz/goinsta/v2"
)

type instagramClient struct {
	cl     *goinsta.Instagram
	photos []Photo
}
type Photo struct {
	url    string
	author string
	likes  int
}

func New(username, password string) *instagramClient {
	insta := goinsta.New(username, password)
	return &instagramClient{cl: insta}
}

func (i *instagramClient) GetHashTagPhotos(name string) error {
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
			i.photos = append(i.photos, Photo{
				author: item.User.FullName,
				url:    item.Images.GetBest(),
				likes:  item.Likes,
			})
		}
		min := 5
		max := 120
		sleepTime := rand.Intn(max-min) + min
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
	return nil
}

func (i *instagramClient) GetPhotos() *[]Photo {
	return &i.photos
}
