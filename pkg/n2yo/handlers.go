package n2yo

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/skrawkipuszczy/discord-bot/pkg/geolocation"
)

const issVisualPassHandlerTxt = "iss pass"

type issVisualPassHandler struct {
	prefix string
	nc     *N2yoClient
}

func NewIssVisualPassHandler(prefix string, c *N2yoClient) *issVisualPassHandler {
	return &issVisualPassHandler{prefix: prefix, nc: c}
}
func (c *issVisualPassHandler) GetCommand() string {
	return fmt.Sprintf("%s %s", c.prefix, issVisualPassHandlerTxt)
}
func (c *issVisualPassHandler) GetDescription() string {
	return "display iss visual passes over the place for next 7 days"
}

func (c *issVisualPassHandler) RegisterDiscordHandler() interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.HasPrefix(m.Content, c.GetCommand()) {
			place := after(m.Content, c.GetCommand())
			location, err := geolocation.GetLocation(place)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("cos poszlo zle %s", err))
			}
			r, err := c.nc.GetISSPass(location)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("cos poszlo zle %s", err))
			}
			var mess string
			for _, m := range r.Passes {
				mess = fmt.Sprintf("%s%s - %s z kierunku od %d do %d\n", mess, startDateFromUnix(m.StartUTC), endDateFromUnix(m.EndUTC), int(m.StartAz), int(m.EndAz))

			}
			ans := &discordgo.MessageEmbed{
				Title:       fmt.Sprintf("Przeloty %s dla %s w ciągu kolejnych 7 dni - ilość %d", r.Info.Satname, place, r.Info.Passescount),
				Description: mess,
				Footer:      &discordgo.MessageEmbedFooter{Text: "Wyincyj klarity kierwa!!!"},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, ans)
		}
	}
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func startDateFromUnix(u int64) string {
	unixTimeUTC := time.Unix(u, 0)
	return unixTimeUTC.Format("2006-01-02 15:04:05")
}
func endDateFromUnix(u int64) string {
	unixTimeUTC := time.Unix(u, 0)

	return unixTimeUTC.Format("15:04:05")
}
