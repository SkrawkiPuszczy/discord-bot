package meteo

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/skrawkipuszczy/discord-bot/pkg/discord"
)

var (
	meteoHandlerCmd = "pogoda"
	footer          = &discordgo.MessageEmbedFooter{Text: "Dane pochodzą z serwisu burze.dzis.net - dziękujemy !!!"}
	author          = &discordgo.MessageEmbedAuthor{Name: "burze.dzis.net"}
)

func (c *meteoClient) RegisterDiscordHandlers() []discord.NamedHandler {
	var handlers []discord.NamedHandler
	handlers = append(handlers, discord.NamedHandler{
		Command:     fmt.Sprintf("%s %s", c.prefix, meteoHandlerCmd),
		Description: "Zwraca warunki pogodowe dla podanej miejscowości",
		Method:      weatherAlertHandler(c),
	})
	return handlers
}

func weatherAlertHandler(c *meteoClient) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		cmd := fmt.Sprintf("%s %s", c.prefix, meteoHandlerCmd)
		if strings.HasPrefix(strings.ToUpper(m.Content), strings.ToUpper(cmd)) {
			place := after(m.Content, cmd)
			d, err := c.fetchCities(place)
			if err != nil {
				log.Println(err)
				return
			}
			ans := createWeatherMesaage(place, d)
			s.ChannelMessageSendEmbed(m.ChannelID, ans)

		}
	}
}
func createWeatherMesaage(place string, d *weather) *discordgo.MessageEmbed {
	var fields []*discordgo.MessageEmbedField
	var description string = ""
	if d.storm.count != 0 {
		azimut := d.storm.azimut
		if azimut == "" {
			azimut = "brak danych"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Dane o wyładowaniach pochodzą z ostatnich (minut)", Value: strconv.Itoa(d.storm.duration)})
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Zaobserwowane wyładowania atomosferyczne:", Value: strconv.Itoa(d.storm.count)})
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Wyładowania w odlegości", Value: floattostr(d.storm.distance)})
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Wyładowania na kierunku", Value: azimut})
	} else {
		description = description + fmt.Sprintln("Brak wyładowań w w tej okolicy")
	}
	if d.alert.fromDate != "" && d.alert.toDate != "" {
		description = fmt.Sprintf("Ostrzeżenia od %s do %s\n", d.alert.fromDate, d.alert.toDate)
		if d.alert.storm == 0 {
			description = description + fmt.Sprintf("Możliwość wystapienia burz od %s do %s\n", d.alert.stormFromDay, d.alert.stormToDay)
		}
	} else {
		description = description + fmt.Sprintln("Brak ostrzeżeń dla podanej okolicy")
	}
	message := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Burze i ostrzeżenia meteo dla %s w promieniu %d km", place, defaultDistance),
		Description: description,
		Author:      author,
		Fields:      fields,
		Footer:      footer,
	}
	if d.storm.count != 0 {
		message.URL = fmt.Sprintf("https://map.blitzortung.org/#9.32/%f/%f", d.lng, d.lat)
	}
	return message
}
func after(value string, a string) string {
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
func floattostr(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}
