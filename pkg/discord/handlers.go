package discord

import (
	"fmt"
	"strings"

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
				Footer:      &discordgo.MessageEmbedFooter{Text: "więcej klarity kierwa!!!"},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, ans)
		}
	}
}
