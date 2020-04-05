package geolocation

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type Location struct {
}

func GetLocation(place string) (*geo.Location, error) {
	return openstreetmap.Geocoder().Geocode(place)
}
