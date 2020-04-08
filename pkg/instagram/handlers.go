package instagram

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/bwmarrin/discordgo"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

const displayRandomInstagramPhotoHandlerTxt = "photo from instagram"

type displayRandomInstagramPhotoHandler struct {
	prefix string
	cache  cache.PhotosCache
}

func NewDisplayRandomInstagramPhotoHandler(prefix string, cache cache.PhotosCache) *displayRandomInstagramPhotoHandler {
	return &displayRandomInstagramPhotoHandler{prefix: prefix, cache: cache}
}
func (c *displayRandomInstagramPhotoHandler) GetCommand() string {
	return fmt.Sprintf("%s %s", c.prefix, displayRandomInstagramPhotoHandlerTxt)
}
func (c *displayRandomInstagramPhotoHandler) GetDescription() string {
	return "display random photo from instagram hashtag"
}

func (c *displayRandomInstagramPhotoHandler) RegisterDiscordHandler() interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if strings.HasPrefix(strings.ToUpper(m.Content), strings.ToUpper(c.GetCommand())) {
			photos, err := c.cache.GetPhotos()
			if err != nil {
				log.Println(err)
				return
			}
			countPhotos := len(photos)
			if countPhotos > 0 {
				t := MapRandomKeyGet(photos).(string)
				var i goinsta.Item
				err = json.Unmarshal([]byte(photos[t]), &i)
				if err != nil {
					log.Println(err)
					return
				}
				photo := &discordgo.MessageEmbedImage{URL: i.Images.GetBest()}
				likes := &discordgo.MessageEmbedField{Name: "likes", Value: strconv.Itoa(i.Likes)}
				ans := &discordgo.MessageEmbed{
					Author: &discordgo.MessageEmbedAuthor{Name: i.User.FullName},
					Image:  photo,
					Fields: []*discordgo.MessageEmbedField{likes},
					Footer: &discordgo.MessageEmbedFooter{Text: "Podoba się - łapka w górę, nie podoba się - łapka w dół"},
				}

				s.ChannelMessageSendEmbed(m.ChannelID, ans)
			}
		}
	}
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}
