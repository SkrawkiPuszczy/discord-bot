package scheduler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jasonlvhit/gocron"
)

func New(configFile string, sFn func(channelID string, message *discordgo.MessageSend)) error {

	gocron.Every(1).Minute().Do(sFn, "694928399910699010", &discordgo.MessageSend{
		Content: "dd",
	})

	return nil
}
