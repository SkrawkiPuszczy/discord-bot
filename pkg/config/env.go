package config

type EnvConfig struct {
	DiscordToken             string `required:"true" split_words:"true"`
	CommandPrefix            string `required:"true" split_words:"true"`
	InstagramEnabled         bool   `split_words:"true" default:"false"`
	InstagramUsername        string `split_words:"true"`
	InstagramPassword        string `split_words:"true"`
	InstagramHashtag         string `split_words:"true"`
	N2yoEnabled              bool   `split_words:"true" default:"false"`
	N2yoApiKey               string `split_words:"true"`
	ScheduledMessagesEnabled bool   `split_words:"true" default:"false"`
	ScheduledConfigFileUrl   string `split_words:"true"`
	CacheType                string `default:"memory" split_words:"true"`
	RedisUrl                 string `split_words:"true"`
	AdMessageInterval        int    `split_words:"true" default:"30"`
}
