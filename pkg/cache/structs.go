package cache

type KeyType string

const (
	Location KeyType = "locations"
	Photo    KeyType = "photos"
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
