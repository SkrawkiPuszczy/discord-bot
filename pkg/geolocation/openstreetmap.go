package geolocation

import (
	"errors"

	"github.com/codingsince1985/geo-golang"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

var (
	UNKNOWN_LOCATION_NAME_ERROR = errors.New("unknown location")
)

func (c *Client) GetLocation(place string) (*geo.Location, error) {
	loc, err := c.c.(cache.LocationCache).GetLocation(place)
	if err != nil {
		return nil, err
	}
	if loc == nil {
		p, err := c.g.Geocode(place)
		if err != nil {
			return nil, err
		}
		if p == nil {
			return nil, UNKNOWN_LOCATION_NAME_ERROR
		}
		loc, err = c.c.(cache.LocationCache).SetLocation(place, p.Lat, p.Lng)
		if err != nil {
			return nil, err
		}
	}
	r := &geo.Location{Lat: loc.Latitude, Lng: loc.Longitude}
	return r, nil
}
