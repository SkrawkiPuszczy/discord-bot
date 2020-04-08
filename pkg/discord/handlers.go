package discord

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type helpHandler struct {
	prefix string
	h      *[]DiscordHandler
}

func NewHelpHandler(prefix string, h *[]DiscordHandler) *helpHandler {
	return &helpHandler{prefix: prefix, h: h}
}
func (c *helpHandler) GetCommand() string {
	return fmt.Sprintf("%s pomoc", c.prefix)
}
func (c *helpHandler) GetDescription() string {
	return "help"
}

func (c *helpHandler) RegisterDiscordHandler() interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.ToUpper(m.Content) == strings.ToUpper(c.GetCommand()) {
			var mess string
			for _, m := range *c.h {
				mess = fmt.Sprintf("%s%s - %s \n", mess, m.GetCommand(), m.GetDescription())

			}
			ans := &discordgo.MessageEmbed{
				Title:       "Dostępne funkcje",
				Description: mess,
				Footer:      &discordgo.MessageEmbedFooter{Text: "Wincyj klarity kierwa!!!"},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, ans)
		}
	}
}

type randomUserHandler struct {
	prefix string
}

func NewRandomUserHandler(prefix string) *randomUserHandler {
	return &randomUserHandler{prefix: prefix}
}
func (c *randomUserHandler) GetCommand() string {
	return fmt.Sprintf("%s wylosuj usera", c.prefix)
}
func (c *randomUserHandler) GetDescription() string {
	return "losuje uźytkownika z serwera"
}

func (c *randomUserHandler) RegisterDiscordHandler() interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.ToUpper(m.Content) == strings.ToUpper(c.GetCommand()) {
			ch, err := s.Channel(m.ChannelID)
			if err != nil {
				log.Println(err)
				return
			}
			users, err := s.GuildMembers(ch.GuildID, "", 1000)
			if err != nil {
				log.Println(err)
				return
			}
			cusers := len(users)
			if cusers > 0 {
				rand.Seed(time.Now().UnixNano())
				user := users[rand.Intn(len(users))]
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Wylosowano %s", user.User.Username))
			}

		}
	}
}
