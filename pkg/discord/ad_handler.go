package discord

import (
	"github.com/bwmarrin/discordgo"
)

type adMessageHandler struct {
	interval int
	c        map[string]int
}

func NewAdMessageHandler(interval int) *adMessageHandler {
	return &adMessageHandler{interval: interval, c: map[string]int{}}
}
func (h *adMessageHandler) AdMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if h.c[m.ChannelID] == h.interval {
		h.c[m.ChannelID] = 0
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
		s.ChannelMessageSendEmbed(m.ChannelID, ans)

	} else {
		h.c[m.ChannelID] = h.c[m.ChannelID] + 1
	}
}
