package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordHandler interface {
	RegisterDiscordHandler() interface{}
	GetCommand() string
	GetDescription() string
}

type discordClient struct {
	s        *discordgo.Session
	handlers []DiscordHandler
}

func New(token string, handlers ...DiscordHandler) (*discordClient, error) {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}
	cl := &discordClient{s: dg}
	for _, c := range handlers {
		cl.handlers = append(cl.handlers, c)
	}
	return cl, nil
}

func (d *discordClient) Run() error {
	for _, c := range d.handlers {
		log.Printf("register discord handler:  %s - %s", c.GetCommand(), c.GetDescription())
		d.s.AddHandler(c.RegisterDiscordHandler())
	}
	return d.s.Open()
}

func (d *discordClient) Close() error {
	return d.s.Close()
}
func (d *discordClient) RegisterHandlers(handlers ...DiscordHandler) {
	for _, c := range handlers {
		d.handlers = append(d.handlers, c)
	}
}
func (d *discordClient) GetHandlers() *[]DiscordHandler {
	return &d.handlers
}
