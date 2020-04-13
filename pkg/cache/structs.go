package cache

import "github.com/skrawkipuszczy/discord-bot/pkg/config"

type KeyType string

const (
	Location KeyType = "locations"
	Photo    KeyType = "photos"
	Weather  KeyType = "weather"
)

type Cache interface {
	Close()
}

type LocationCache interface {
	SetLocation(name string, lat, long float64) (*location, error)
	GetLocation(name string) (*location, error)
}
type PhotosCache interface {
	SetPhoto(keyName string, data string) error
	GetPhotos() (map[string]string, error)
}
type MessagesOnChannelsCache interface {
	SetChannelMessagesCounter(keyName string, data int) error
	GetChannelMessagesCounter(keyName string) (int, error)
}
type WetherCitiesCache interface {
	SetWetherCityLocation(keyName string, x, y float64) error
	GetWetherCityLocation(keyName string) (float64, float64, error)
}
type CacheReader interface {
	Get(s KeyType, keyName string) ([]byte, error)
}

type CacheWriter interface {
	Set(s KeyType, keyName string, data []byte) ([]byte, error)
}

type location struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func New(config *config.EnvConfig) (Cache, error) {
	if config.CacheType == "redis" {
		return NewRedisClient(config.RedisURL)
	} else {
		return NewMemoryCache()
	}
}
