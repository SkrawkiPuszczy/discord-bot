package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	DiscordToken             string `required:"true" split_words:"true"`
	CommandPrefix            string `required:"true" split_words:"true"`
	InstagramEnabled         bool   `split_words:"true" default:"false"`
	InstagramUsername        string `split_words:"true"`
	InstagramPassword        string `split_words:"true"`
	InstagramHashtag         string `split_words:"true"`
	N2yoEnabled              bool   `split_words:"true" default:"false"`
	N2yoAPIKey               string `split_words:"true"`
	ScheduledMessagesEnabled bool   `split_words:"true" default:"false"`
	ScheduledConfigFileURL   string `split_words:"true"`
	CacheType                string `default:"memory" split_words:"true"`
	RedisURL                 string `split_words:"true"`
	AdMessageInterval        int    `split_words:"true" default:"30"`
	DatabaseURL              string `split_words:"true" required:"true"`
	BurzeDzisAPIKey          string `split_words:"true"`
	HTTPAdmin                string `split_words:"true" default:"admin"`
	HTTPPassword             string `split_words:"true" default:"password"`
	HTMLStaticDir            string `split_words:"true" default:"./web/"`
}

func New() (*EnvConfig, error) {
	var c EnvConfig
	err := envconfig.Process("bot", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
