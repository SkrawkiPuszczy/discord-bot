package cache

import (
	"encoding/json"
)

type memoryCache struct {
	locations map[string][]byte
	photos    map[string]string
}

func NewMemoryCache() (*memoryCache, error) {
	return &memoryCache{
		locations: map[string][]byte{},
		photos:    map[string]string{},
	}, nil
}
func (c *memoryCache) Close() {
	return
}

//SetLocation set geo coordinates for place in cache
func (r *memoryCache) SetLocation(name string, lat, long float64) (*location, error) {
	data := &location{Name: name, Latitude: lat, Longitude: long}
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r.locations[prepareKey(Location, name)] = json
	return data, nil
}

//GetLocation return geo coordinates for place
func (r *memoryCache) GetLocation(name string) (*location, error) {
	str := r.locations[prepareKey(Location, name)]
	loc := &location{}
	if len(str) == 0 {
		return nil, nil
	}
	err := json.Unmarshal([]byte(str), &loc)
	if err != nil {
		return nil, err
	}
	return loc, nil

}
func (r *memoryCache) SetPhoto(keyName string, data string) error {
	r.photos[keyName] = data
	return nil
}
func (r *memoryCache) GetPhotos() (map[string]string, error) {
	return r.photos, nil
}
