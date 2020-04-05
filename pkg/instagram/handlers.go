package instagram

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const displayRandomInstagramPhotoHandlerTxt = "photo from instagram"

type displayRandomInstagramPhotoHandler struct {
	prefix string
	photos *[]Photo
}

func NewDisplayRandomInstagramPhotoHandler(prefix string, photos *[]Photo) *displayRandomInstagramPhotoHandler {
	return &displayRandomInstagramPhotoHandler{prefix: prefix, photos: photos}
}
func (c *displayRandomInstagramPhotoHandler) GetCommand() string {
	return fmt.Sprintf("%s %s", c.prefix, displayRandomInstagramPhotoHandlerTxt)
}
func (c *displayRandomInstagramPhotoHandler) GetDescription() string {
	return "display random photo from instagram hashtag"
}

func (c *displayRandomInstagramPhotoHandler) RegisterDiscordHandler() interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if strings.HasPrefix(m.Content, c.GetCommand()) {
			photos := *c.photos
			countPhotos := len(photos)
			log.Printf("photos %d", countPhotos)
			if countPhotos > 0 {
				rand.Seed(time.Now().UnixNano())
				choosenPhoto := photos[rand.Intn(countPhotos)]
				photo := &discordgo.MessageEmbedImage{URL: choosenPhoto.url}
				likes := &discordgo.MessageEmbedField{Name: "likes", Value: strconv.Itoa(choosenPhoto.likes)}
				ans := &discordgo.MessageEmbed{
					Author: &discordgo.MessageEmbedAuthor{Name: choosenPhoto.author},
					Image:  photo,
					Fields: []*discordgo.MessageEmbedField{likes},
				}

				s.ChannelMessageSendEmbed(m.ChannelID, ans)
			}
		}
	}
}
