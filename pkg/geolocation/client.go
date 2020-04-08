package geolocation

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

type LocationReader interface {
	GetLocation(place string) (*geo.Location, error)
}

type Client struct {
	g geo.Geocoder
	c cache.Cache
}

func New(c cache.Cache) *Client {
	return &Client{
		g: openstreetmap.Geocoder(),
		c: c,
	}
}
