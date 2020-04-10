package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v7"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

type adMessageHandler struct {
	interval int
	c        cache.MessagesOnChannelsCache
}

func NewAdMessageHandler(c cache.MessagesOnChannelsCache, interval int) *adMessageHandler {
	return &adMessageHandler{interval: interval, c: c}
}
func (h *adMessageHandler) AdMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	count, err := h.c.GetChannelMessagesCounter(m.ChannelID)
	if err == redis.Nil {
		count = 0
	} else if err != nil {
		log.Println(err)
		return
	}
	if count == h.interval {
		mess := `
		Zostań patronem Skrawki Puszczy i pomóż rozwijać kanał

		Patronite: https://patronite.pl/SkrawkiPuszczy
		Tipeo: https://www.tipeo.pl/SkrawkiPuszczy

		Zapraszam również na zakupy w sklepie: https://skrawkipuszczy.sklep.pl/
		`
		ans := &discordgo.MessageEmbed{
			Title:       "Jak możesz wspomóc działalność Skrawki Puszczy",
			Description: mess,
			Footer:      &discordgo.MessageEmbedFooter{Text: "Wyincyj klarity kierwa!!!"},
		}

		h.c.SetChannelMessagesCounter(m.ChannelID, 0)
		s.ChannelMessageSendEmbed(m.ChannelID, ans)

	} else {
		count = count + 1
		h.c.SetChannelMessagesCounter(m.ChannelID, count)
	}
}
