package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type MessageType int

const (
	SimpleMessage MessageType = iota
	MessageEmbed
)

type DiscordHandler interface {
	RegisterDiscordHandler() interface{}
	GetCommand() string
	GetDescription() string
}

type NamedHandler struct {
	Command     string
	Description string
	Method      interface{}
}
type DiscordHandlerProvider interface {
	RegisterDiscordHandlers() []NamedHandler
	// AddDiscordHandlers() []struct {
	// 	command     string
	// 	description string
	// 	funC        interface{}
	// }
}
type discordClient struct {
	s        *discordgo.Session
	handlers []DiscordHandler
	h        []DiscordHandlerProvider
}

func New(token string, handlers ...DiscordHandler) (*discordClient, error) {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}
	cl := &discordClient{s: dg, h: []DiscordHandlerProvider{}}
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
	for _, c := range d.h {
		for _, w := range c.RegisterDiscordHandlers() {
			log.Printf("register discord handler:  %s - %s", w.Command, w.Description)
			d.s.AddHandler(w.Method)
		}

	}
	return d.s.Open()
}

func (d *discordClient) Close() error {
	log.Println("discord client closed")
	return d.s.Close()
}

func (d *discordClient) RegisterHandlers(handlers ...DiscordHandler) {
	for _, c := range handlers {
		d.handlers = append(d.handlers, c)
	}
}
func (d *discordClient) RegisterProvidedHandlers(handlers ...DiscordHandlerProvider) {
	for _, c := range handlers {
		d.h = append(d.h, c)
	}
}
func (d *discordClient) GetHandlers() *[]DiscordHandler {
	return &d.handlers
}
func (d *discordClient) GetSession() *discordgo.Session {
	return d.s
}

func (d *discordClient) SendMessage(channelID string, message *discordgo.MessageSend) {
	c, _ := d.s.Channel(channelID)
	log.Printf("send scheduled message to channel: %s \n", c.Name)
	d.s.ChannelMessageSend(channelID, message.Content)
}
