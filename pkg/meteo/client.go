package meteo

import (
	"fmt"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/skrawkipuszczy/discord-bot/pkg/cache"
)

var (
	defaultDistance = 50
)

type meteoClient struct {
	bCl    SerwerSOAPPort
	db     gorm.DB
	cache  cache.WetherCitiesCache
	apiKey string
	prefix string
}

func New(apiKey string, db gorm.DB, cache cache.WetherCitiesCache, prefix string) *meteoClient {
	soapCli := &soap.Client{
		URL:       "https://burze.dzis.net/soap.php",
		Namespace: Namespace,
	}
	burze := NewSerwerSOAPPort(soapCli)
	burze.KeyAPI(apiKey)
	db.AutoMigrate(City{})
	return &meteoClient{bCl: burze, db: db, apiKey: apiKey, cache: cache, prefix: prefix}
}

type stormData struct {
	duration int
	count    int
	distance float64
	azimut   string
}
type alertData struct {
	fromDate       string
	toDate         string
	storm          int
	stormToDay     string
	stormFromDay   string
	tornado        int
	tornadoToDay   string
	tornadoFromDay string
	wind           int
	windToDay      string
	windFromDay    string
	rain           int
	rainToDay      string
	rainFromDay    string
}

type weather struct {
	storm *stormData
	alert *alertData
	lat   float64
	lng   float64
}

func (c *meteoClient) fetchCities(city string) (*weather, error) {
	lat, lng, err := c.cache.GetWetherCityLocation(city)

	if err != nil {
		return nil, errors.Wrap(err, "fetch cities cache")
	}
	if lat == 1000 && lng == 1000 {
		s, err := c.bCl.Miejscowosc(city, c.apiKey)
		if err != nil {
			return nil, errors.Wrap(err, "fetch cities from burze.dzis.net")
		}
		lat = *s.X
		lng = *s.Y
		err = c.cache.SetWetherCityLocation(city, lat, lng)
		if err != nil {
			return nil, errors.Wrap(err, "fetch cities from burze.dzis.net")
		}
	}
	alert, err := c.fetchAlertsData(lng, lat)
	if err != nil {
		return nil, errors.Wrap(err, "fetch weather alerts from burze.dzis.net")
	}
	storm, err := c.fetchStormData(lng, lat)
	if err != nil {
		return nil, errors.Wrap(err, "fetch storm data from burze.dzis.net")
	}
	return &weather{
		lat:   lat,
		lng:   lng,
		alert: alert,
		storm: storm,
	}, nil
}

func (c *meteoClient) fetchAlertsData(lng, lat float64) (*alertData, error) {
	d, err := c.bCl.Ostrzezenia_pogodowe(lng, lat, c.apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "fetch wether alerts from burze.dzis.net")
	}
	alert := &alertData{
		fromDate:       *d.Od_dnia,
		toDate:         *d.Do_dnia,
		storm:          *d.Burza,
		stormFromDay:   *d.Burza_od_dnia,
		stormToDay:     *d.Burza_do_dnia,
		tornado:        *d.Traba,
		tornadoFromDay: *d.Traba_od_dnia,
		tornadoToDay:   *d.Traba_do_dnia,
		wind:           *d.Wiatr,
		windFromDay:    *d.Wiatr_od_dnia,
		windToDay:      *d.Wiatr_do_dnia,
		rain:           *d.Opad,
		rainFromDay:    *d.Opad_od_dnia,
		rainToDay:      *d.Opad_do_dnia,
	}
	return alert, nil
}

func (c *meteoClient) fetchStormData(lng, lat float64) (*stormData, error) {
	b, err := c.bCl.Szukaj_burzy(fmt.Sprintf("%v", lng), fmt.Sprintf("%v", lat), defaultDistance, c.apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "fetch wether alerts from burze.dzis.net")
	}
	storm := &stormData{
		duration: *b.Okres,
		count:    *b.Liczba,
		distance: *b.Odleglosc,
		azimut:   *b.Kierunek,
	}
	return storm, nil
}
